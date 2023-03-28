package routes

import (
	"gin-gorm-clean-template/controller"
	"gin-gorm-clean-template/service"

	"github.com/gin-gonic/gin"
)

func FeedsRoutes(router *gin.Engine, FeedsController controller.FeedsController, jwtService service.JWTService) {
	FeedsRoutes := router.Group("/api/feeds")
	{
		FeedsRoutes.GET("", FeedsController.GetAllFeeds)
	}
}