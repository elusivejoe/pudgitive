package main

//TODO: main.go exists for development purposes; to be removed.

import (
	"os"

	"github.com/elusivejoe/pudgitive/database"
	"github.com/elusivejoe/pudgitive/wrapper"
)

func main() {
	db := database.NewDatabase("tmp/testdb")
	wrapper := wrapper.NewWrapper(db)
	wrapper.InitRoot("test")
	db.Close()

	os.RemoveAll("tmp")
}
