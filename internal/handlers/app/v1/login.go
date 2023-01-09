package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xybor-x/xyerror"
	"github.com/xybor/xyauth/pkg/service"
)

type LoginParams struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func LoginGETHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}

func LoginPOSTHandler(ctx *gin.Context) {
	params := new(LoginParams)
	ctx.ShouldBind(params)

	err := service.Authenticate(params.Email, params.Password)

	if err != nil {
		switch {
		case errors.Is(err, service.NotFoundError):
			ctx.HTML(http.StatusNotFound, "login.html", gin.H{"message": xyerror.Message(err)})
		case errors.Is(err, service.CredentialError):
			ctx.HTML(http.StatusForbidden, "login.html", gin.H{"message": xyerror.Message(err)})
		case err != nil:
			ctx.HTML(http.StatusInternalServerError, "login.html", gin.H{
				"message": "Something is wrong, please contact to administrator if the issue persists",
			})
		}
		return
	}

	token, err := service.CreateToken(params.Email)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "login.html", gin.H{
			"message": "Something is wrong, please contact to administrator if the issue persists",
		})
		return
	}

	ctx.SetCookie("access_token", token, 5*60, "/", "localhost", true, true)
	ctx.Redirect(http.StatusMovedPermanently, "")
}
