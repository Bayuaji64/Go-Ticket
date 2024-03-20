package models

import (
	"errors"
	"fmt"
	"log"
	"time"

	"example.com/rest-api/db"
)

type CheckoutDetail struct {
	ID       int64
	User     UserDetails
	ShowTime ShowTimeDetails
	Seat     SeatDetails
}

type UserDetails struct {
	ID    int64
	Name  string
	Email string
}

type ShowTimeDetails struct {
	ID       int64
	DateTime time.Time
	Price    float64
	Movie    MovieDetails
}

type MovieDetails struct {
	ID    int64
	Title string
}

type SeatDetails struct {
	ID         int64
	SeatNumber string
	Status     SeatStatus
}

type CancelBookingRequest struct {
	CheckoutID int64
}

type PaymentRequest struct {
	CheckoutID int64
	Amount     float64
}

func BookSeat(userID, showtimeID, seatID int64) (int64, error) {
	tx, err := db.DB.Begin()
	if err != nil {
		return 0, err
	}

	if _, err = tx.Exec("UPDATE seats SET status = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $1", seatID, 2); err != nil {
		tx.Rollback()
		return 0, err
	}

	var checkoutID int64

	err = tx.QueryRow("INSERT INTO checkouts (user_id, showtime_id, seat_id) VALUES ($1, $2, $3) RETURNING id", userID, showtimeID, seatID).Scan(&checkoutID)
	if err != nil {
		tx.Rollback()
		log.Println("Failed to insert into checkouts:", err)
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return checkoutID, nil
}
func CancelBooking(checkoutId int64) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}

	var seatId int64
	if err := tx.QueryRow("SELECT seat_id FROM checkouts WHERE id = $1", checkoutId).Scan(&seatId); err != nil {
		tx.Rollback()
		return fmt.Errorf("could not find checkout with id %d: %v", checkoutId, err)
	}

	if _, err = tx.Exec("UPDATE seats SET status = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $1", seatId, 2); err != nil {
		tx.Rollback()
		return err
	}

	if _, err = tx.Exec("DELETE FROM checkouts WHERE id = $1", checkoutId); err != nil {
		tx.Rollback()
		return fmt.Errorf("could not delete checkout: %v", err)
	}

	return tx.Commit()
}

func GetCheckoutDetail(checkoutId int64) (*CheckoutDetail, error) {
	var detail CheckoutDetail
	query := `
SELECT c.id, u.id, u.name, u.email, s.id, s.date_time, s.price, m.id, m.title, seat.id, seat.seat_number, seat.status
FROM checkouts c
JOIN users u ON c.user_id = u.id
JOIN showtimes s ON c.showtime_id = s.id
JOIN movies m ON s.movie_id = m.id
JOIN seats seat ON c.seat_id = seat.id
WHERE c.id = $1`

	err := db.DB.QueryRow(query, checkoutId).Scan(&detail.ID, &detail.User.ID, &detail.User.Name, &detail.User.Email, &detail.ShowTime.ID, &detail.ShowTime.DateTime, &detail.ShowTime.Price, &detail.ShowTime.Movie.ID, &detail.ShowTime.Movie.Title, &detail.Seat.ID, &detail.Seat.SeatNumber, &detail.Seat.Status)
	if err != nil {
		log.Println("Error querying checkout detail:", err)
		return nil, err
	}

	return &detail, nil
}

func ProcessPayment(checkoutId int64, amount float64) error {
	var price, seatID float64

	err := db.DB.QueryRow(`
        SELECT c.id, s.price, c.seat_id FROM checkouts c
        JOIN showtimes s ON c.showtime_id = s.id
        WHERE c.id = $1
    `, checkoutId).Scan(&checkoutId, &price, &seatID)
	if err != nil {
		log.Println("Error retrieving booking details:", err)
		return err
	}

	if amount != price {
		return errors.New("the amount paid does not match")
	}

	_, err = db.DB.Exec("UPDATE seats SET status = 3 WHERE id = $1", seatID)
	if err != nil {
		log.Println("Error updating seat status:", err)
		return err
	}

	return nil
}
