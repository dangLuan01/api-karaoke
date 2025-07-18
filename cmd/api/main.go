package main

import (
	"log"

	"github.com/dangLuan01/karaoke/internal/app"
	"github.com/dangLuan01/karaoke/internal/config"
	"github.com/dangLuan01/karaoke/internal/db"
)

func main() {
	
	app.LoadEnv()

	if err := db.InitDB(); err != nil {
		log.Fatalf("unable to connect to sql")
	}
	
	cfg := config.NewConfig()

	application := app.NewApplication(cfg, db.DB)

	if err := application.Run(); err != nil {
		panic(err)
	}
}