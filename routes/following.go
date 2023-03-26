package routes

import (
	"gin-gorm-clean-template/controller"
	"gin-gorm-clean-template/middleware"
	"gin-gorm-clean-template/service"

	"github.com/gin-gonic/gin"
)

func FollowingRoutes(router *gin.Engine, FollowingController controller.FollowingController, jwtService service.JWTService) {
	followingRoutes := router.Group("/api/following")
	{
		followingRoutes.POST("", middleware.CreateShortUrlAuthenticate(jwtService, false), FollowingController.CreateFollowing)
		followingRoutes.GET("/me", middleware.CreateShortUrlAuthenticate(jwtService, false), FollowingController.FindFollowingByUserID)
	}
}