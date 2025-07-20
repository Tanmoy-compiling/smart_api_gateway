package user

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all the user service routes
func SetupRoutes(router *gin.Engine) {
	// User routes group
	userGroup := router.Group("/api/users")
	{
		// Authentication routes
		userGroup.POST("/register", RegisterHandler)
		userGroup.POST("/login", LoginHandler)
		userGroup.POST("/logout", AuthMiddleware(), LogoutHandler)

		// User profile routes
		userGroup.GET("/profile", AuthMiddleware(), GetProfileHandler)
		userGroup.PUT("/profile", AuthMiddleware(), UpdateProfileHandler)
		
		// Password management
		userGroup.POST("/forgot-password", ForgotPasswordHandler)
		userGroup.POST("/reset-password", ResetPasswordHandler)
		
		// Admin routes
		adminGroup := userGroup.Group("/admin")
		adminGroup.Use(AuthMiddleware(), AdminMiddleware())
		{
			adminGroup.GET("/all", GetAllUsersHandler)
			adminGroup.DELETE("/:id", DeleteUserHandler)
			adminGroup.PUT("/:id/status", UpdateUserStatusHandler)
		}
	}
}

// Handler function declarations
func RegisterHandler(c *gin.Context)
func LoginHandler(c *gin.Context)
func LogoutHandler(c *gin.Context)
func GetProfileHandler(c *gin.Context)
func UpdateProfileHandler(c *gin.Context)
func ForgotPasswordHandler(c *gin.Context)
func ResetPasswordHandler(c *gin.Context)
func GetAllUsersHandler(c *gin.Context)
func DeleteUserHandler(c *gin.Context)
func UpdateUserStatusHandler(c *gin.Context)

// Middleware declarations
func AuthMiddleware() gin.HandlerFunc
func AdminMiddleware() gin.HandlerFunc
