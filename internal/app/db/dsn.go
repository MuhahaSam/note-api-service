package db

type DbConfig struct {
	Host     string
	User     string
	Password string
	DbName   string
}

func GetFakeConfig() *DbConfig {
	return &DbConfig{
		Host:     "postgres.host",
		User:     "user",
		Password: "password",
		DbName:   "dbName",
	}
}
