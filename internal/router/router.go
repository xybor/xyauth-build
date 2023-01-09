package router

import (
	"github.com/gin-gonic/gin"
	apiv1 "github.com/xybor/xyauth/internal/handlers/api/v1"
	appv1 "github.com/xybor/xyauth/internal/handlers/app/v1"
	"github.com/xybor/xyauth/internal/handlers/well_known/openid"
	"github.com/xybor/xyauth/internal/middlewares"
)

func New() *gin.Engine {
	router := gin.Default()

	router.Static("/static", "web/static")
	router.StaticFile("/favicon.ico", "web/static/favicon.ico")
	router.LoadHTMLGlob("web/template/*.html")

	router.GET(".well-known/openid-configuration", openid.Handler)
	router.GET("login", appv1.LoginGETHandler)
	router.POST("login", appv1.LoginPOSTHandler)
	router.GET("register", appv1.RegisterGETHandler)
	router.POST("register", appv1.RegisterPOSTHandler)

	mustAuthGroup := router.Group("")
	mustAuthGroup.Use(middlewares.ValidateAccessToken)
	{
		mustAuthGroup.GET("", appv1.WelcomeHandler)
	}

	apiv1Group := router.Group("api/v1")
	{
		apiv1Group.POST("register", apiv1.RegisterHandler)
	}

	return router
}
