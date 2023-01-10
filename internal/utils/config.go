package utils

import (
	"github.com/xybor-x/xyconfig"
)

var config *xyconfig.Config

func initConfig() {
	config = xyconfig.GetConfig("xyauth")

	if err := config.ReadFile("configs/default.ini", true); err != nil {
		panic(err)
	}

	if config.GetDefault("general.environment", "dev").MustString() == "dev" {
		if err := config.ReadFile(".env", true); err != nil {
			panic(err)
		}
	}

	d := config.GetDefault("general.env_watch_cycle", 0).MustDuration()
	if err := config.LoadEnv(d); err != nil {
		panic(err)
	}
}

func GetConfig() *xyconfig.Config {
	return config
}
