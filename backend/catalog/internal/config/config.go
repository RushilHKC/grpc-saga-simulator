package config

import (
	"os"
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
		Addr: os.Getenv("ADDR"),
	}
}

func Load_Db_Config() *DB_config {
	return &DB_config{
		Port:     os.Getenv("port"),
		User:     os.Getenv("user"),
		Password: os.Getenv("password"),
		Db_name:  os.Getenv("db_name"),
	}
}
