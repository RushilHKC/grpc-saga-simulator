package config

type DB_config struct {
	Port         int
	User         string
	Password     string
	Db_name      string
	MaxIdleConns int
	MaxOpenConn  int
}

func Load_Db_Config() *DB_config {
	return &DB_config{
		Port:     5432,
		User:     "Harshwardhan",
		Password: "admin",
		Db_name:  "catalog",
	}
}
