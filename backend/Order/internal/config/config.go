package config

import "github.com/RushilHKC/grpcSimulator/backend/Order/internal/env"

type serverConfig struct {
	Addr string
}

type DbConfig struct {
	Port         string
	UserName     string
	Password     string
	DbName       string
	MaxIdleConns int
	MaxOpenConns int
	DbTimeOut    int
}

func LoadServerConfig() *serverConfig {
	addr := env.GetString("Addr")

	return &serverConfig{Addr: addr}
}

func LoadDbConfig() *DbConfig {
	port := env.GetString("port")
	userName := env.GetString("userName")
	password := env.GetString("password")
	dbName := env.GetString("dbName")
	maxIdleConns := env.GetInt("maxIdleConns")
	maxOpenConns := env.GetInt("maxOpenConns")
	dbTimeOut := env.GetInt("database_time_out")

	return &DbConfig{
		Port:         port,
		UserName:     userName,
		Password:     password,
		DbName:       dbName,
		MaxIdleConns: maxIdleConns,
		MaxOpenConns: maxOpenConns,
		DbTimeOut:    dbTimeOut,
	}
}
