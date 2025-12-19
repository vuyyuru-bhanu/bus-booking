package routes

import (
	"github.com/gin-gonic/gin"
	"bus-booking/controllers"
	"bus-booking/middleware"
)

func SetupRoutes(r *gin.Engine) {
	// Group all API routes under /api
	api := r.Group("/api")
	api.Use(middleware.CORSMiddleware())

	// Public routes
	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)
	api.GET("/notifications", controllers.GetNotifications)
	api.GET("/routes", controllers.GetRoutes)

	// Protected routes (require user authentication)
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/buses", controllers.GetBuses)
		protected.POST("/bookings", controllers.CreateBooking)

		protected.GET("/profile", controllers.GetProfile)
		protected.PUT("/profile", controllers.UpdateProfile)

		protected.POST("/payments", controllers.InitiatePayment)
		protected.GET("/payments/:payment_id", controllers.GetPaymentStatus)

		protected.GET("/bookings/history", controllers.GetBookingHistory)
		protected.GET("/bookings/history/:id", controllers.GetBookingDetails)

		protected.GET("/buses/:bus_id/seats", controllers.GetAvailableSeats)
		protected.POST("/buses/:bus_id/seats", controllers.SelectSeats)

		protected.GET("/buses/:bus_id/reviews", controllers.GetReviews)
		protected.POST("/buses/:bus_id/reviews", controllers.AddReview)
	}

	// Admin routes (require admin role)
	admin := api.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.POST("/routes", controllers.CreateRoute)
		admin.POST("/buses", controllers.CreateBus)
		admin.POST("/buses/:bus_id/image", controllers.UploadBusImage)
		admin.POST("/notifications", controllers.CreateNotification)
		admin.GET("/bookings", controllers.GetAllBookings)

		admin.GET("/users", controllers.GetUsers)
		admin.POST("/users", controllers.CreateUser)
		admin.PUT("/users/:id", controllers.UpdateUser)
		admin.DELETE("/users/:id", controllers.DeleteUser)

		admin.PUT("/bookings/:id/cancel", controllers.CancelBooking)
	}

	// Serve static files (bus images)
	r.Static("/uploads", "./uploads")
}
