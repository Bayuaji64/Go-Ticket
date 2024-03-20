package models

import (
	"encoding/json"
	"log"
	"time"

	"example.com/rest-api/db"
)

type MovieDetail struct {
	ID       int64
	Movie    string
	DateTime time.Time
	Seats    []Seat
}
type Seat struct {
	ID         int64
	ShowtimeID int64
	SeatNumber string
	Status     SeatStatus
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type SeatStatus int

const (
	StatusAvailable SeatStatus = 1
	StatusPending   SeatStatus = 2
	StatusCompleted SeatStatus = 3
)

type BookingRequest struct {
	UserID int64
	SeatID int64
}

func (s SeatStatus) MarshalJSON() ([]byte, error) {
	var statusStr string
	switch s {
	case StatusAvailable:
		statusStr = "available"
	case StatusPending:
		statusStr = "pending"
	case StatusCompleted:
		statusStr = "completed"
	default:
		statusStr = "unknown"
	}
	return json.Marshal(statusStr)
}

func GetMovieDetailWithShowtimeAndSeats(movieID, showtimeID int64) (*MovieDetail, error) {

	movieQuery := `SELECT id, title FROM movies WHERE id = $1`
	var movieTitle string
	var id int64
	err := db.DB.QueryRow(movieQuery, movieID).Scan(&id, &movieTitle)
	if err != nil {
		log.Println("Error querying movie by ID:", err)
		return nil, err
	}

	showtimeQuery := `SELECT date_time FROM showtimes WHERE id = $1 AND movie_id = $2`
	var dateTime time.Time
	err = db.DB.QueryRow(showtimeQuery, showtimeID, movieID).Scan(&dateTime)
	if err != nil {
		log.Println("Error querying showtime by ID:", err)
		return nil, err
	}

	seatsQuery := `SELECT id, showtime_id, seat_number, status FROM seats WHERE showtime_id = $1 AND status = 1`
	rows, err := db.DB.Query(seatsQuery, showtimeID)
	if err != nil {
		log.Println("Error querying seats for showtime:", err)
		return nil, err
	}
	defer rows.Close()

	var seats []Seat
	for rows.Next() {
		var seat Seat
		err = rows.Scan(&seat.ID, &seat.ShowtimeID, &seat.SeatNumber, &seat.Status)
		if err != nil {
			log.Println("Error scanning seat:", err)
			return nil, err
		}
		seats = append(seats, seat)
	}

	detail := &MovieDetail{
		ID:       id,
		Movie:    movieTitle,
		DateTime: dateTime,
		Seats:    seats,
	}

	return detail, nil
}
