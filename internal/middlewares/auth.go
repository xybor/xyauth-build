package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xybor/xyauth/pkg/token"
)

func ValidateAccessToken(ctx *gin.Context) {
	accessToken, err := ctx.Cookie("access_token")
	if err != nil {
		ctx.Redirect(http.StatusTemporaryRedirect, "login")
		return
	}

	_, err = token.ValidateUAT(accessToken)
	if err != nil {
		ctx.Redirect(http.StatusTemporaryRedirect, "login")
		return
	}

	ctx.Next()
}
