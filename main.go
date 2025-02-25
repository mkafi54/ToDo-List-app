package main

import (
	"time"

	"github.com/kahfi/to-do-list-app/config"
	"github.com/kahfi/to-do-list-app/database"
	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("SERVER_TIMEZONE", "Asia/Jakarta")
	loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
	time.Local = loc

	if err := config.SetupConfig(); err != nil {
		// logger.Fatalf("config SetupConfig() error: %s", err)
	}
	lmsDSN, replicaDSNLms := config.DbConfiguration("prototype-lms")

	if err := database.DbConnection(lmsDSN, replicaDSNLms); err != nil {
		// logger.Fatalf("database DbConnection error: %s", err)
	}
}
