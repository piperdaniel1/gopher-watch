package main

import (
	"fmt"
	"log"
	"os"

	"github.com/piperdaniel1/gopher-watch/server/config"
	"github.com/piperdaniel1/gopher-watch/server/db"
	"github.com/piperdaniel1/gopher-watch/server/models"
)

func main() {
	file, err := os.Open("config.json")

	if err != nil {
		log.Fatalf("Failed to open configuration file: %v", err)
	}
	defer file.Close()

	cfg, err := config.Load(file)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	conn, err := db.GetDBConnection(cfg)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatalf("Error with connection!")
	} else {
		fmt.Println("Connection is validated and successful.")
	}

	// add a server
	// server, err := models.CreateServer(&models.Server{
	// 	ComputerName: sql.NullString{Valid: true, String: "ArchSystem"},
	// 	OS:           sql.NullString{Valid: true, String: "Arch Linux"},
	// }, conn)
	// if err != nil {
	// 	log.Fatalf("Could not create server: %v", err)
	// }
	// log.Printf("Added new server: %+v", server)

	// get a server
	servers, err := models.GetAllServers(conn)
	if err != nil {
		log.Fatalf("Could not get servers: %v", err)
	}
	log.Printf("Got servers: %+v", servers)
}
