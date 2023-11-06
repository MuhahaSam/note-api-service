package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	dbPassEscSeq = "{password}"
)

// DB ...
type DB struct {
	DSN                string `json:"dsn"`
	MaxOpenConnections int32  `json:"max_open_connections"`
}

// NoteConfig ...
type NoteConfig struct {
	GrpcConf *GrpcConf   `json:"grpc"`
	HttpConf *HttpConf   `json:"http"`
	Logger   *LoggerConf `json:"logger"`
	DB       *DB         `json:"db"`
}

type LoggerConf struct {
	ShowTime bool `json:"show_time"`
}

type GrpcConf struct {
	Host string `json: "host"`
	Port string `json: "port"`
}

type HttpConf struct {
	Host string `json: "host"`
	Port string `json: "port"`
}

// NewNoteConfig ...
func NewNoteConfig(path string) (*NoteConfig, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &NoteConfig{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// GetLoggerConfig ...
func (c *NoteConfig) GetGrpcHost() *GrpcConf {
	return c.GrpcConf
}

// GetSourceConfig ...
func (c *NoteConfig) GetHttpHost() *HttpConf {
	return c.HttpConf
}

func (c *NoteConfig) GetLoggerConfig() *LoggerConf {
	return c.Logger
}

// GetDBConfig ...
func (c *NoteConfig) GetDBConfig() (*pgxpool.Config, error) {
	dbDsn := strings.ReplaceAll(c.DB.DSN, dbPassEscSeq, os.Getenv("DB_PASSWORD"))
	fmt.Println(dbDsn)

	poolConfig, err := pgxpool.ParseConfig(dbDsn)
	if err != nil {
		return nil, err
	}
	poolConfig.ConnConfig.BuildStatementCache = nil
	poolConfig.ConnConfig.PreferSimpleProtocol = true
	poolConfig.MaxConns = c.DB.MaxOpenConnections

	return poolConfig, nil
}
