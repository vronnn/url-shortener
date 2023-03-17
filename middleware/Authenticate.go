package middleware

import (
	"fmt"
	"gin-gorm-clean-template/common"
	"gin-gorm-clean-template/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(jwtService service.JWTService, isAdmin bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Ditemukan", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		if !strings.Contains(authHeader, "Bearer ") {
			response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
		token, err := jwtService.ValidateToken(authHeader)
		if err != nil {
			response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		if !token.Valid {
			response := common.BuildErrorResponse("Gagal Memproses Request", "Akses Ditolak", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if isAdmin {
			userRole, err := jwtService.GetUserRoleByToken(authHeader)
			if err != nil {
				response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}
			fmt.Println(userRole)
			if userRole != "admin" {
				response := common.BuildErrorResponse("Gagal Memproses Request", "Role User Tidak Memiliki Akses ke Endpoint Ini", nil)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}
		}

		userID, err := jwtService.GetUserIDByToken(authHeader)
		if err != nil {
			response := common.BuildErrorResponse("Gagal Memproses Request", err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		ctx.Set("token", authHeader)
		ctx.Set("userID", userID)
		ctx.Next()
	}
}