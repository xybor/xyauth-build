package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/xybor-x/xyerror"
	"github.com/xybor/xyauth/internal/database"
	"github.com/xybor/xyauth/internal/router"
	"github.com/xybor/xyauth/internal/utils"
)

var configs []string

var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Long:  "Start the authentication and authorization server",
	Run: func(cmd *cobra.Command, args []string) {
		for i := range configs {
			if err := utils.GetConfig().ReadFile(configs[i], true); err != nil {
				panic(err)
			}
		}

		database.Initialize()
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
			panic(xyerror.BaseException.Newf("invalid environment %s", env))
		}

		host := utils.GetConfig().GetDefault("server.host", "localhost").MustString()
		port := utils.GetConfig().GetDefault("server.port", 443).MustInt()
		addr := fmt.Sprintf("%s:%d", host, port)

		privateKey := utils.GetConfig().MustGet("server.private_key_path").MustString()
		publicKey := utils.GetConfig().MustGet("server.public_key_path").MustString()

		utils.GetLogger().Event("server-start").Field("address", addr).Info()
		utils.GetLogger().Event("server-close").
			Field("error", http.ListenAndServeTLS(addr, publicKey, privateKey, r)).Info()
	},
}

func main() {
	rootCmd.Flags().StringArrayVarP(&configs, "config", "c", nil, "Specify overridden config files")
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
