package config

import (
	"github.com/RushilHKC/grpcSimulator/backend/catalog/internal/env"
)

type DB_config struct {
	Port         string
	User         string
	Password     string
	Db_name      string
	MaxIdleConns int
	MaxOpenConn  int
}

type Server_config struct {
	Addr string
}

func LoadServerConfig() *Server_config {
	return &Server_config{
		Addr: env.GetString("ADDR"),
	}
}

func Load_Db_Config() *DB_config {
	return &DB_config{
		Port:     env.GetString("port"),
		User:     env.GetString("user"),
		Password: env.GetString("password"),
		Db_name:  env.GetString("db_name"),
	}
}
