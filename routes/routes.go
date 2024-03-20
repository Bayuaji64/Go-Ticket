package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func AllRoutes(server *gin.Engine) {

	server.POST("/signup", register)
	server.POST("/login", login)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.GET("/movies", handleMovies)
	authenticated.POST("/movies", PostBookSeat)
	authenticated.DELETE("/movies", CancelBooking)
	authenticated.GET("/movies-checkout/:checkoutId", GetCheckoutDetailHandler)

	authenticated.PUT("/movies-payment", ProcessPaymentHandler)

}
