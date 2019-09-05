package main

import (
	"github.com/BurntSushi/toml"
	"github.com/maddevsio/comedian/api"
	"github.com/maddevsio/comedian/config"
	"github.com/maddevsio/comedian/storage"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	log "github.com/sirupsen/logrus"
	"golang.org/x/text/language"
)

func main() {

	cnf, err := config.Get()
	if err != nil {
		log.Fatal("Failed to get config : ", err)
	}

	db, err := storage.New(cnf.DatabaseURL, "migrations")
	if err != nil {
		log.Fatal("Failed to connect to db: ", err)
	}

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("active.en.toml")
	bundle.MustLoadMessageFile("active.ru.toml")

	comedian := api.New(cnf, db, bundle)

	if err = comedian.Start(); err != nil {
		log.Fatal("Failed to start Comedian API: ", err)
	}
}
