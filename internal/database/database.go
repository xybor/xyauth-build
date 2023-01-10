package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/xybor/xyauth/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var config = utils.GetConfig()

func Initialize() {
	dsnFormat := "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s"
	// Prioritize environment variable over config file.
	host := config.MustGet("database.host").MustString()
	port := config.MustGet("database.port").MustString()
	user := config.MustGet("POSTGRES_USER").MustString()
	password := config.MustGet("POSTGRES_PASSWORD").MustString()
	dbname := config.MustGet("POSTGRES_DB").MustString()
	sslmode := config.GetDefault("database.sslmode", "disable").MustString()
	dsn := fmt.Sprintf(dsnFormat, host, user, password, dbname, port, sslmode)

	timezone, ok := config.Get("database.timezone")
	if ok {
		dsn += fmt.Sprintf(" TimeZone=%s", timezone.MustString())
	}

	loglevel := config.GetDefault("database.loglevel", logger.Info).MustInt()
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.LogLevel(loglevel),
			IgnoreRecordNotFoundError: true,
		},
	)

	var err error
	nRetries := config.GetDefault("database.retries", 3).MustInt()
	retryDuration := config.GetDefault("database.retry_duration", time.Second).MustDuration()
	for i := 0; i < nRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
		if err == nil {
			break
		}
		utils.GetLogger().Event("connect-to-database-failed").Field("error", err).Warning()
		time.Sleep(retryDuration)
	}

	utils.GetLogger().Event("connect-do-database").
		Field("host", host).Field("port", port).Info()

	if err != nil {
		panic(err)
	}
}

func Get() *gorm.DB {
	return db
}
