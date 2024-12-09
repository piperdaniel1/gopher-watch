package config

import (
	"encoding/json"
	"io"
)

type Config struct {
	DBHost     string `json:"DBHost"`
	DBPort     int    `json:"DBPort"`
	DBUser     string `json:"DBUser"`
	DBPassword string `json:"DBPassword"`
	DBName     string `json:"DBName"`
}

type DBConfig interface {
	GetDBHost() string
	GetDBPort() int
	GetDBUser() string
	GetDBPassword() string
	GetDBName() string
}

func Load(reader io.Reader) (*Config, error) {
	var cfg Config
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (cfg *Config) GetDBHost() string     { return cfg.DBHost }
func (cfg *Config) GetDBPort() int        { return cfg.DBPort }
func (cfg *Config) GetDBUser() string     { return cfg.DBUser }
func (cfg *Config) GetDBPassword() string { return cfg.DBPassword }
func (cfg *Config) GetDBName() string     { return cfg.DBName }
