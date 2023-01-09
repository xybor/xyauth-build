package v1

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xybor/xyauth/pkg/token"
)

func WelcomeHandler(ctx *gin.Context) {
	accessToken, err := ctx.Cookie("access_token")
	if err != nil {
		ctx.Redirect(http.StatusTemporaryRedirect, "login")
		return
	}

	info, err := token.ValidateUAT(accessToken)
	if err != nil {
		ctx.Redirect(http.StatusTemporaryRedirect, "login")
		return
	}

	name := strings.Split(info.Email, "@")[0]
	ctx.HTML(http.StatusOK, "welcome.html", gin.H{"name": name})
}
