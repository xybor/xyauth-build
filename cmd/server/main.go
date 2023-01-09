package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xybor-x/xyerror"
	"github.com/xybor/xyauth/internal/router"
	"github.com/xybor/xyauth/internal/utils"
)

func main() {
	r := router.New()
	env := utils.GetConfig().GetDefault("general.environment", "dev").MustString()
	switch env {
	case "dev":
		gin.SetMode(gin.DebugMode)
	case "test", "staging":
		gin.SetMode(gin.TestMode)
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	default:
		panic(xyerror.BaseException.Newf("Invalid environment %s", env))
	}

	host := utils.GetConfig().GetDefault("server.host", "localhost").MustString()
	port := utils.GetConfig().GetDefault("server.port", 443).MustInt()
	addr := fmt.Sprintf("%s:%d", host, port)

	privateKey := utils.GetConfig().MustGet("server.private_key_path").MustString()
	publicKey := utils.GetConfig().MustGet("server.public_key_path").MustString()

	utils.GetLogger().Event("server-start").Field("address", addr).Info()
	utils.GetLogger().Event("server-close").
		Field("error", http.ListenAndServeTLS(addr, publicKey, privateKey, r)).Info()
}
