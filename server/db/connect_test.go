package db

import (
	"log"
	"os"
	"testing"

	"github.com/piperdaniel1/gopher-watch/server/config"
)

func TestMain(m *testing.M) {
	// Load the test database
	file, err := os.Open("config_test.json")
	if err != nil {
		log.Fatalf("Unable to open test config: %v", err)
	}

	cfg, err := config.Load(file)
	if err != nil {
		log.Fatalf("Unable to open load config: %v", err)
	}

	conn, err := GetDBConnection(cfg)
	if err != nil {
		log.Fatalf("Unable to get DB connection: %v", err)
	}
	defer conn.Close()

	// Execute reset script
	resetScript, err := os.ReadFile("reset_db.sql")
	if err != nil {
		log.Fatalf("Unable to open reset database script: %v", err)
	}

	_, err = conn.Exec(string(resetScript))
	if err != nil {
		log.Fatalf("Error while executing database reset script: %v", err)
	}

	code := m.Run()

	os.Exit(code)
}

// integration test
// uses real config module, connects to the real database
func TestGetDBConnection(t *testing.T) {
	file, err := os.Open("config_test.json")
	if err != nil {
		t.Fatalf("Unable to open test config: %v", err)
	}
	cfg, err := config.Load(file)
	if err != nil {
		t.Fatalf("Unable to open load config: %v", err)
	}
	conn, err := GetDBConnection(cfg)

	if err != nil {
		t.Fatalf("Failed to get connection with error: %v", err)
	}

	if conn == nil {
		t.Fatalf("Received nil connection variable.")
	}

	if conn.Ping() != nil {
		t.Fatalf("Could not ping returned database connection.")
	}

	err = conn.Close()

	if err != nil {
		t.Fatalf("Received error when attempting to close returned connection: %v", err)
	}
}
