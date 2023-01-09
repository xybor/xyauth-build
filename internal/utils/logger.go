package utils

import (
	"fmt"
	"os"

	"github.com/xybor-x/xyconfig"
	"github.com/xybor-x/xylog"
)

var lg *xylog.Logger

func initLogger() {
	emitter := xylog.NewStreamEmitter(os.Stdout)
	handler := xylog.GetHandler("xybor.auth")
	handler.AddMacro("time", "asctime")
	handler.AddMacro("level", "levelname")
	handler.AddEmitter(emitter)

	lg = xylog.GetLogger("xybor.auth")
	lg.SetLevel(config.GetDefault("general.loglevel", xylog.INFO).MustInt())
	lg.AddHandler(handler)

	config.AddHook("general.loglevel", func(e xyconfig.Event) {
		lg.SetLevel(e.New.MustInt())
		fmt.Println("Set log level to ", e.New.MustInt())
	})
}

func GetLogger() *xylog.Logger {
	return lg
}
