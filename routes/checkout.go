package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

type PaymentRequest struct {
	CheckoutID int64
	Amount     float64
}

func PostBookSeat(c *gin.Context) {

	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userIDInt, ok := userID.(int64) // Pastikan tipe userID sesuai dengan yang diharapkan
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	movieID, err := strconv.ParseInt(c.Query("movieID"), 10, 64)
	if err != nil || movieID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing movieID"})
		return
	}

	showtimeID, err := strconv.ParseInt(c.Query("showtimeID"), 10, 64)
	if err != nil || showtimeID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing showtimeID"})
		return
	}

	var req BookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	resultChan := make(chan struct {
		checkoutID int64
		err        error
	})
	go func() {
		checkoutID, err := models.BookSeat(userIDInt, showtimeID, req.SeatID)
		resultChan <- struct {
			checkoutID int64
			err        error
		}{checkoutID, err}
	}()

	result := <-resultChan
	if result.err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Booking failed", "detail": result.err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Booking successful",
		"checkoutID": result.checkoutID,
		"ShowTimeID": showtimeID,
		"SeatID":     req.SeatID,
	})
}

func CancelBooking(c *gin.Context) {
	var req models.CancelBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := models.CancelBooking(req.CheckoutID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cancellation failed", "detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking cancelled successfully"})
}

func GetCheckoutDetailHandler(c *gin.Context) {
	checkoutId, err := strconv.ParseInt(c.Param("checkoutId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid checkout ID"})
		return
	}

	detail, err := models.GetCheckoutDetail(checkoutId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve checkout details"})
		return
	}

	c.JSON(http.StatusOK, detail)
}

func ProcessPaymentHandler(c *gin.Context) {
	var req PaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	resultChan := make(chan error)
	go func() {
		err := models.ProcessPayment(req.CheckoutID, req.Amount)
		resultChan <- err
	}()

	err := <-resultChan
	if err != nil {
		if err.Error() == "the amount paid does not match" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Payment processing failed", "detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment processed successfully"})
}
