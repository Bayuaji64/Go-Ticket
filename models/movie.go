package models

import (
	"log"
	"time"

	"example.com/rest-api/db"
)

type Movie struct {
	ID        int64
	Title     string
	Genre     string
	CreatedAt time.Time
	UpdatedAt time.Time
	ShowTimes []ShowTime `json:"showTimes,omitempty"`
}

func GetAllMovies() ([]Movie, error) {

	query := `SELECT * FROM movies`
	rows, err := db.DB.Query(query)
	if err != nil {
		log.Println("Error querying for movies:", err)
		return nil, err
	}
	defer rows.Close()

	var movies []Movie
	for rows.Next() {
		var m Movie
		err := rows.Scan(&m.ID, &m.Title, &m.Genre, &m.CreatedAt, &m.UpdatedAt)
		if err != nil {
			log.Println("Error :", err)
			return nil, err
		}
		movies = append(movies, m)
	}
	return movies, nil
}

func DetailMovieByID(id int64) (*Movie, error) {
	movieQuery := `SELECT id, title, genre, created_at, updated_at FROM movies WHERE id = $1`
	showtimeQuery := `SELECT id,movie_id ,date_time, price FROM showtimes WHERE movie_id = $1 ORDER BY date_time`

	var m Movie
	row := db.DB.QueryRow(movieQuery, id)
	err := row.Scan(&m.ID, &m.Title, &m.Genre, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		log.Println("Error querying movie by ID:", err)
		return nil, err
	}

	rows, err := db.DB.Query(showtimeQuery, id)
	if err != nil {
		log.Println("Error querying showtimes for movie:", err)
		return nil, err
	}
	defer rows.Close()

	var showtimes []ShowTime
	for rows.Next() {
		var st ShowTime
		if err := rows.Scan(&st.ID, &st.MovieID, &st.DateTime, &st.Price); err != nil {
			log.Println("Error scanning showtime:", err)
			return nil, err
		}
		showtimes = append(showtimes, st)
	}

	m.ShowTimes = showtimes

	return &m, nil
}
