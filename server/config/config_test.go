package config

import (
	"strings"
	"testing"
)

func TestLoad(t *testing.T) {
	jsonConfig := `{
		"DBHost": "localhost",
		"DBPort": 5432,
		"DBUser": "root",
		"DBPassword": "password",
		"DBName": "gopher_watch"
	}`

	reader := strings.NewReader(jsonConfig)

	cfg, err := Load(reader)
	if err != nil {
		t.Fatalf("Expected no error, got error on configuration load: %v", err)
	}

	if cfg.DBHost != "localhost" {
		t.Errorf("Expedted DBHost to be 'localhost', got '%s'", cfg.DBHost)
	}

	if cfg.DBPort != 5432 {
		t.Errorf("Expedted DBPort to be '5432', got '%d'", cfg.DBPort)
	}
	if cfg.DBUser != "root" {
		t.Errorf("Expedted DBUser to be 'root', got '%s'", cfg.DBUser)
	}
	if cfg.DBPassword != "password" {
		t.Errorf("Expedted DBPassword to be 'password', got '%s'", cfg.DBPassword)
	}
	if cfg.DBName != "gopher_watch" {
		t.Errorf("Expedted DBName to be 'gopher_watch', got '%s'", cfg.DBName)
	}
}
