package main

import (
    "github.com/gin-gonic/gin"
    "smart_api_gateway/pkg/common/database"
    "smart_api_gateway/pkg/common/middleware"
    "smart_api_gateway/services/user/handlers"
)

func main() {
    r := gin.Default()
    
    // Initialize DB
    if err := database.InitDB(); err != nil {
        panic(err)
    }
    
    // Routes
    v1 := r.Group("/api/v1")
    {
        v1.POST("/register", handlers.Register)
        v1.POST("/login", handlers.Login)
        v1.GET("/profile", middleware.AuthRequired(), handlers.GetProfile)
    }
    
    r.Run(":5001")
}