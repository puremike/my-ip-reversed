package main

import (
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
	"github.com/puremike/ip_reversed/internal/db"
	"github.com/puremike/ip_reversed/internal/store"
	"github.com/puremike/ip_reversed/util"
)

const PORT = ":8180"

type application struct {
	AppConfig *AppConfig
	Store     *store.Storage
}

type AppConfig struct {
	PORT    string
	Db_addr string
}

func main() {

	var db_addr string

	env := strings.TrimSpace(os.Getenv("APP_ENV"))
	if env == "development" {
		db_addr = os.Getenv("DB_ADDR")
	} else {
		db_addr = db.GetPostgresDSN()
	}

	cfg := &AppConfig{
		Db_addr: db_addr,
		PORT:    util.GetEnvString("PORT", PORT),
	}

	db, err := db.NewPostgresDB(cfg.Db_addr)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	defer db.Close()
	log.Print("Connected to database successfully")

	app := &application{
		AppConfig: cfg,
		Store:     store.NewStorage(db),
	}

	mux := app.route()

	if err := app.server(mux); err != nil {
		panic(err)
	}
}
