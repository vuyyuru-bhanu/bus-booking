package routes

import (
	"github.com/gin-gonic/gin"
	"bus-booking/controllers"
	"bus-booking/middleware"
)

func SetupRoutes(r *gin.Engine) {
	// Public routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/notifications", controllers.GetNotifications)

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/buses", controllers.GetBuses)
		protected.POST("/bookings", controllers.CreateBooking)
	}

	// Admin routes
	admin := r.Group("/api/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.POST("/routes", controllers.CreateRoute)
		admin.POST("/buses", controllers.CreateBus)
		admin.POST("/buses/:bus_id/image", controllers.UploadBusImage)
		admin.POST("/notifications", controllers.CreateNotification)
		admin.GET("/bookings", controllers.GetAllBookings)
	}

	// Serve static files (bus images)
	r.Static("/uploads", "./uploads")
}