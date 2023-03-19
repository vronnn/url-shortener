package routes

import (
	"gin-gorm-clean-template/controller"
	"gin-gorm-clean-template/middleware"
	"gin-gorm-clean-template/service"

	"github.com/gin-gonic/gin"
)

func UrlShortenerRoutes(router *gin.Engine, UrlShortenerController controller.UrlShortenerController, jwtService service.JWTService) {
	urlShortenerRoutes := router.Group("/api/url_shortener")
	{
		urlShortenerRoutes.POST("", middleware.CreateShortUrlAuthenticate(jwtService, false), UrlShortenerController.CreateUrlShortener)
	}
}