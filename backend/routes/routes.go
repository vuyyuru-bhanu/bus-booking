package routes

import (
	"github.com/gin-gonic/gin"
	"bus-booking/controllers"
	"bus-booking/middleware"
)

func SetupRoutes(r *gin.Engine) {
	// Group all API routes under /api
	api := r.Group("/api")

	// Public routes
	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)
	api.GET("/notifications", controllers.GetNotifications)

	// Protected routes (require user authentication)
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/buses", controllers.GetBuses)
		protected.POST("/bookings", controllers.CreateBooking)
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
	}

	// Serve static files (bus images)
	r.Static("/uploads", "./uploads")
}
