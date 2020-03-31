package main

import (
	"github.com/SungminSo/qr-generator/models"
	"github.com/SungminSo/qr-generator/server"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.AutomaticEnv()

	// Set mongoDB
	models.InitDB(viper.GetString("mongo_db_host"), viper.GetString("mongo_db_port"))

	bindAddr := viper.GetString("BIND_ADDR")
	if bindAddr == "" {
		bindAddr = "0.0.0.0:3506"
	}

	s := server.NewServer(bindAddr)
	log.Fatal(s.Run())
}
