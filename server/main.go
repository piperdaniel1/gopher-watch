package main

import (
	"fmt"

	"github.com/piperdaniel1/gopher-watch/server/db"
)

func main() {
	db.GetDBConnection()
	fmt.Println("This is the server package.")
}
