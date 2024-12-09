package models

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateServerAdd(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Could not create mock database: %v", err)
	}

	mock.ExpectQuery(regexp.QuoteMeta(`
        INSERT INTO servers (computer_name, os)
        VALUES ($1, $2)
        RETURNING id;
    `)).
		WithArgs(sql.NullString{String: "TestComputerName", Valid: true}, sql.NullString{String: "TestOS", Valid: true}).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	_, err = CreateServer(&Server{
		ComputerName: sql.NullString{Valid: true, String: "TestComputerName"},
		OS:           sql.NullString{Valid: true, String: "TestOS"},
	}, db)

	if err != nil {
		t.Fatalf("Error on creating a server: %v", err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatalf("Failed to call database correctly: %v", err)
	}
}

func TestCreateServerErrorOnRowNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Could not create mock database: %v", err)
	}
	mock.ExpectQuery(regexp.QuoteMeta(`
        INSERT INTO servers (computer_name, os)
        VALUES ($1, $2)
        RETURNING id;
    `)).
		WithArgs(sql.NullString{String: "TestComputerName", Valid: true}, sql.NullString{String: "TestOS", Valid: true})

	_, err = CreateServer(&Server{
		ComputerName: sql.NullString{Valid: true, String: "TestComputerName"},
		OS:           sql.NullString{Valid: true, String: "TestOS"},
	}, db)

	if err == nil {
		t.Fatalf("Failed to throw row not found error.")
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatalf("Failed to call database correctly: %v", err)
	}
}

func TestGetServer(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Could not create mock database: %v", err)
	}
	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT id, computer_name, os
		FROM servers
		WHERE id=$1
    `)).
		WithArgs(2).
		WillReturnRows(sqlmock.NewRows([]string{"id", "computer_name", "os"}).AddRow(2, "ArchSystem", sql.NullString{Valid: false}))

	server, err := GetServer(2, db)

	if err != nil {
		t.Fatalf("Error on getting a server: %v", err)
	}

	if server.ID != 2 {
		t.Fatalf("Mismatched server ID")
	}
	if server.ComputerName.String != "ArchSystem" {
		t.Fatalf("Mismatched computer name")
	}
	if server.OS.Valid != false {
		t.Fatalf("Mismatched server ID")
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatalf("Failed to call database correctly: %v", err)
	}
}

func TestGetNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Could not create mock database: %v", err)
	}
	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT id, computer_name, os
		FROM servers
		WHERE id=$1
    `)).
		WithArgs(2)

	_, err = GetServer(2, db)

	if err == nil {
		t.Fatalf("Did not return an error")
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatalf("Failed to call database correctly: %v", err)
	}
}
