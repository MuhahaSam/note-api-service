package db

type DbConfig struct {
	Host     string
	User     string
	Password string
	DbName   string
}

var fakeDbConfig = DbConfig{
	Host:     "postgres.host",
	User:     "user",
	Password: "password",
	DbName:   "dbName",
}
