package controller

import (
	"gin-gorm-clean-template/common"
	"gin-gorm-clean-template/dto"
	"gin-gorm-clean-template/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FollowingController interface {
	CreateFollowing(ctx *gin.Context)
	FindFollowingByUserID(ctx *gin.Context)
}

type followingController struct {
	followingService service.FollowingService
	jwtService service.JWTService
}

func NewFollowingController(fs service.FollowingService, js service.JWTService) FollowingController {
	return &followingController{
		followingService: fs,
		jwtService: js,
	}
}

func(fc *followingController) CreateFollowing(ctx *gin.Context) {
	var followingDTO dto.CreateFollowingDTO
	err := ctx.ShouldBind(&followingDTO)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Following", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	token := ctx.MustGet("token").(string)
	userID, err := fc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	followingDTO.UserID = userID

	checkDuplicate := fc.followingService.CheckDuplicate(ctx, userID.String(), followingDTO.FollowingID.String())
	if !checkDuplicate {
		response := common.BuildErrorResponse("Gagal Menambahkan Following", "User Sudah Follow User Ini", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	result, err := fc.followingService.CreateFollowing(ctx.Request.Context(), followingDTO)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Following", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Menambahkan Following", result)
	ctx.JSON(http.StatusOK, res)
}

func(fc *followingController) FindFollowingByUserID(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := fc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	result, err := fc.followingService.FindFollowingByUserID(ctx.Request.Context(), userID.String())
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Following", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Mendapatkan List Following", result)
	ctx.JSON(http.StatusOK, res)
}