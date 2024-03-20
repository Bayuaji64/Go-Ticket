package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

type BookingRequest struct {
	UserID int64
	SeatID int64
}

func handleMovies(context *gin.Context) {
	movieID := context.Query("movieID")
	showtimeID := context.Query("showtimeID")

	if movieID != "" && showtimeID != "" {
		getShowTimeSeatMovieDetails(context)
	} else if movieID != "" {
		getByIDMovie(context)
	} else {
		getMovies(context)
	}
}

func getMovies(context *gin.Context) {

	events, err := models.GetAllMovies()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, events)

}

func getByIDMovie(context *gin.Context) {
	movieID, err := strconv.ParseInt(context.Query("movieID"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movieID parameter"})
		return
	}

	movie, err := models.DetailMovieByID(movieID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch movie"})
		return
	}

	context.JSON(http.StatusOK, movie)
}

func getShowTimeSeatMovieDetails(context *gin.Context) {
	movieID, err := strconv.ParseInt(context.Query("movieID"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movieID"})
		return
	}

	showtimeID, err := strconv.ParseInt(context.Query("showtimeID"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid showtimeID"})
		return
	}

	details, err := models.GetMovieDetailWithShowtimeAndSeats(movieID, showtimeID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch movie details"})
		return
	}

	context.JSON(http.StatusOK, details)
}
