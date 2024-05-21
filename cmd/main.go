package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"

	"github.com/ShipIM/go-group-manager/internal/postgres"

	"github.com/spf13/viper"
)

type Config struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DBName       string
	SSLMode      string
	MigrationDir string
}

//go:embed migration/*.sql
var MigrationsFS embed.FS

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("can not initialize configs: %s", err.Error())
	}

	cfg := Config{
		Host:         viper.GetString("db.host"),
		Port:         viper.GetString("db.port"),
		Username:     viper.GetString("db.username"),
		Password:     viper.GetString("db.password"),
		DBName:       viper.GetString("db.dbname"),
		SSLMode:      viper.GetString("db.sslmode"),
		MigrationDir: viper.GetString("db.migration_dir"),
	}

	migrator := postgres.GetNewMigrator(MigrationsFS, cfg.MigrationDir)

	dataSource := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)

	conn, err := sql.Open("postgres", dataSource)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	err = migrator.ApplyMigrations(conn)
	if err != nil {
		panic(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("application")

	return viper.ReadInConfig()
}
