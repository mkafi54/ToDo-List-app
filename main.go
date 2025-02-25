package main

import (
	"fmt"
	"time"

	"github.com/kahfi/to-do-list-app/config"
	"github.com/kahfi/to-do-list-app/database"
	"github.com/kahfi/to-do-list-app/routes"
	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("SERVER_TIMEZONE", "Asia/Jakarta")
	loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
	time.Local = loc

	if err := config.SetupConfig(); err != nil {
		fmt.Printf("config SetupConfig() error: %s", err)
	}
	lmsDSN, replicaDSNLms := config.DbConfiguration("postgres")

	if err := database.DbConnection(lmsDSN, replicaDSNLms); err != nil {
		fmt.Printf("database DbConnection error: %s", err)
	}

	router := routes.SetupRoute()
	fmt.Printf("%v", router.Run(config.ServerConfig()))
}
