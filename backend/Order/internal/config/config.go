package config

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
}

func LoadServerConfig() *serverConfig {
	addr := env.GetString("Addr")

	return &serverConfig{Addr: addr}
}

func loadDbConfig() *DbConfig {
	port := env.GetInt("port")
	userName := env.GetString("userName")
	password := env.GetString("password")
	dbName := env.GetString("dbName")
	maxIdleConns := env.getInt("maxIdleConns")
	maxOpenConns := env.getInt("maxOpenConns")

	return &DbConfig{
		Port:         port,
		UserName:     userName,
		Password:     password,
		DbName:       dbName,
		MaxIdleConns: maxIdleConns,
		MaxOpenConns: maxOpenConns,
	}
}
