package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"

	"github.com/ShipIM/go-group-manager/internal/postgres"

	"github.com/spf13/viper"
)

//go:embed migration/*.sql
var MigrationsFS embed.FS

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("can not initialize configs: %s", err.Error())
	}

	var (
		host         = viper.GetString("db.host")
		port         = viper.GetString("db.port")
		username     = viper.GetString("db.username")
		password     = viper.GetString("db.password")
		dbName       = viper.GetString("db.dbname")
		sslMode      = viper.GetString("db.sslmode")
		migrationDir = viper.GetString("db.migration_dir")
	)

	migrator := postgres.NewMigrator(MigrationsFS, migrationDir)

	dataSource := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, port, username, dbName, password, sslMode)

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
