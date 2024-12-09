package models

import (
	"database/sql"
	"fmt"

	"github.com/piperdaniel1/gopher-watch/server/db"
)

type Server struct {
	ID           int            `json:"id"`
	ComputerName sql.NullString `json:"computer_name"`
	OS           sql.NullString `json:"os"`
}

func CreateServer(server *Server, db db.DBInterface) (*Server, error) {
	query := `
		INSERT INTO servers (computer_name, os)
		VALUES ($1, $2)
		RETURNING id;
	`
	err := db.QueryRow(query, server.ComputerName, server.OS).Scan(&server.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to insert server: %w", err)
	}

	return server, nil
}

func GetServer(serverID int, db db.DBInterface) (*Server, error) {
	query := `
		SELECT id, computer_name, os
		FROM servers
		WHERE id=$1
	`

	server := Server{}

	err := db.QueryRow(query, serverID).Scan(&server.ID, &server.ComputerName, &server.OS)

	if err != nil {
		return nil, fmt.Errorf("failed to get server: %v", err)
	}

	return &server, nil
}

func GetAllServers(db db.DBInterface) (*[]Server, error) {
	query := `
		SELECT id, computer_name, os
		FROM servers
	`

	var servers []Server

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var nextServer Server
		if err := rows.Scan(&nextServer.ID, &nextServer.ComputerName, &nextServer.OS); err != nil {
			return &servers, err
		}
		servers = append(servers, nextServer)
	}

	return &servers, nil
}
