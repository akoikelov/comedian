package storage

import (

	// This line is must for working MySQL database
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/maddevsio/comedian/config"
)

func setupDB() *DB {
	c, err := config.Get()
	if err != nil {
		log.Fatal(err)
	}
	db, err := New(c)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
