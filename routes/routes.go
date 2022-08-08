package routes

import (
	"fmt"
	"mitramas_test/controllers"
	"mitramas_test/db"
	"mitramas_test/middlewares"
	"mitramas_test/repositories"
	"mitramas_test/services"
	"os"

	"github.com/gin-gonic/gin"
)

func Init() {
	db := db.Connect()
	defer db.Close()

	// AUTH
	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	// CHECK
	activityRepository := repositories.NewActivityRepository(db)
	activityService := services.NewActivityService(activityRepository)
	activityController := controllers.NewActivityController(activityService)

	gin.SetMode(gin.ReleaseMode)
	e := gin.New()
	e.Use(gin.Recovery(), middlewares.Logger(), middlewares.CORSMiddleware())
	v1 := e.Group("/api/v1")
	{
		v1.POST("/register", authController.RegisterController)
		v1.POST("/login", authController.LoginController)
		v1.POST("/logout", authController.LogoutController).Use(middlewares.Auth())
	}

	activity := v1.Group("check-in").Use(middlewares.Auth())
	{
		activity.POST("", activityController.CheckInController)
		activity.POST("/activity", activityController.CreateActivityController)
		activity.GET("/activity", activityController.ReadActivityController)
		activity.PUT("/activity", activityController.UpdateActivityController)
		activity.DELETE("/activity", activityController.DeleteActivityController)
	}

	attendance := v1.Group("attendance").Use(middlewares.Auth())
	{
		attendance.GET("", activityController.ReadAttendanceController)
	}

	checkOut := v1.Group("check-out").Use(middlewares.Auth())
	{
		checkOut.POST("", activityController.CheckOutController)
	}

	e.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}
