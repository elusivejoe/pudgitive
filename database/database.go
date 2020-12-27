package database

import (
	"log"

	"github.com/recoilme/pudge"
)

func NewDatabase() *pudge.Db {
	cfg := pudge.DefaultConfig
	cfg.SyncInterval = 0

	db, err := pudge.Open("tmp/testdb", cfg)

	if err != nil {
		log.Fatal("Cannot open the database")
	}

	return db
}
