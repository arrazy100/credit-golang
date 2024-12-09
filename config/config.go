package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Server struct {
	AllowOrigins     []string `yaml:"allow_origin"`
	AllowMethods     []string `yaml:"allow_methods"`
	AllowHeaders     []string `yaml:"allow_headers"`
	ExposeHeaders    []string `yaml:"expose_headers"`
	AllowContentType []string `yaml:"allow_content_type"`
}

type Db struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"db_name"`
}

type App struct {
	Mode string `yaml:"mode"`
	Port string `yaml:"port"`
}

type YamlConfig struct {
	Server   Server `yaml:"server"`
	Database Db     `yaml:"db"`
	App      App    `yaml:"app"`
}

type Config struct {
	ServerConfig       Server
	DatabaseConnection *gorm.DB
	App                App
}

func Load(filename string) (*Config, error) {
	config_file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var yamlConfig YamlConfig
	err = yaml.Unmarshal(config_file, &yamlConfig)
	if err != nil {
		return nil, err
	}

	ServerConfig(&yamlConfig.Server)

	db_conn, err := DatabaseConfig(yamlConfig.Database)
	if err != nil {
		return nil, err
	}

	AppConfig(&yamlConfig.App)

	return &Config{
		ServerConfig:       yamlConfig.Server,
		DatabaseConnection: db_conn,
		App:                yamlConfig.App,
	}, nil
}

func ServerConfig(server *Server) {
	if len(server.AllowOrigins) == 0 {
		server.AllowOrigins = []string{"*"}
	}

	if len(server.AllowMethods) == 0 {
		server.AllowMethods = []string{}
	}
}

func DatabaseConfig(database Db) (*gorm.DB, error) {
	if database.Host == "" || database.User == "" || database.Password == "" || database.DbName == "" || database.Port == 0 {
		return nil, errors.New("database config is not complete")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		database.Host, database.User, database.Password, database.DbName, database.Port,
	)
	db_conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db_conn.DB()
	if err != nil {
		return nil, err
	}

	if err = sqlDB.Ping(); err != nil {
		return nil, err
	}

	return db_conn, nil
}

func AppConfig(app *App) {
	if app.Mode == "" {
		app.Mode = "debug"
	}

	if app.Port == "" {
		app.Port = "8080"
	}
}
