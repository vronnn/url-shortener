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
		urlShortenerRoutes.GET("", UrlShortenerController.GetAllUrlShortener)
		urlShortenerRoutes.GET("/:id", UrlShortenerController.GetUrlShortenerByID)
		urlShortenerRoutes.GET("/user/:id", UrlShortenerController.GetUrlShortenerByUserID)
		urlShortenerRoutes.GET("/long_url/:short_url", UrlShortenerController.GetUrlShortenerByShortUrl)
		urlShortenerRoutes.PUT("/:id", middleware.CreateShortUrlAuthenticate(jwtService, false), UrlShortenerController.UpdateUrlShortener)
		urlShortenerRoutes.DELETE("/:id", middleware.CreateShortUrlAuthenticate(jwtService, false), UrlShortenerController.DeleteUrlShortener)
	}
}