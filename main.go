package main

//TODO: main.go exists for development purposes; to be removed.

import (
	"fmt"
	"log"

	"github.com/elusivejoe/pudgitive/database"
)

func main() {
	db := database.NewDatabase()

	db.Set("test_key", []string{"test", "value", "one"})

	var out []string

	if err := db.Get("test_key", &out); err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Printf("%v", out)
}
