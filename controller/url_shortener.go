package controller

import (
	"gin-gorm-clean-template/common"
	"gin-gorm-clean-template/dto"
	"gin-gorm-clean-template/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UrlShortenerController interface {
	CreateUrlShortener(ctx *gin.Context)
}

type urlShortenerController struct {
	urlShortenerService service.UrlShortenerService
	jwtService service.JWTService
}

func NewUrlShortenerController(us service.UrlShortenerService, js service.JWTService) UrlShortenerController {
	return &urlShortenerController{
		urlShortenerService: us,
		jwtService: js,
	}
}

func(uc *urlShortenerController) CreateUrlShortener(ctx *gin.Context) {
	var urlShortener dto.UrlShortenerCreateDTO
	err := ctx.ShouldBind(&urlShortener)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Url Shortener", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if ctx.Request.Header["Authorization"] != nil {
		token := ctx.MustGet("token").(string)
		userID, err := uc.jwtService.GetUserIDByToken(token)
		if err != nil {
			response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		urlShortener.UserID = userID
	}
	
	
	checkUrlShortener, _ := uc.urlShortenerService.GetUrlShortenerByShortUrl(ctx.Request.Context(), urlShortener.ShortUrl)
	if checkUrlShortener.ShortUrl != "" {
		res := common.BuildErrorResponse("Gagal Menambahkan Url Shortener", "Short Url Sudah Terdaftar", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if *urlShortener.Private && urlShortener.Password == "" {
		res := common.BuildErrorResponse("Gagal Menambahkan Url Shortener", "Url Shortener Private Harus Mengandung Password", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := uc.urlShortenerService.CreateUrlShortener(ctx.Request.Context(), urlShortener)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Url Shortener", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Menambahkan Url Shortener", result)
	ctx.JSON(http.StatusOK, res)
}