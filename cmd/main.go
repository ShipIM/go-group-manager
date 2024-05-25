package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"

	"github.com/ShipIM/go-group-manager/api/rest/handler"
	"github.com/ShipIM/go-group-manager/internal/postgres"
	"github.com/ShipIM/go-group-manager/internal/repository"
	"github.com/ShipIM/go-group-manager/internal/service"
	"github.com/ShipIM/go-group-manager/pkg"

	"github.com/spf13/viper"
)

//go:embed migration/*.sql
var MigrationsFS embed.FS

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("can not initialize configs: %s", err.Error())
	}

	var (
		serverPort   = viper.GetString("server.port")
		host         = viper.GetString("db.host")
		dbPort       = viper.GetString("db.port")
		username     = viper.GetString("db.username")
		password     = viper.GetString("db.password")
		dbName       = viper.GetString("db.dbname")
		sslMode      = viper.GetString("db.sslmode")
		migrationDir = viper.GetString("db.migration_dir")
	)

	migrator := postgres.NewMigrator(MigrationsFS, migrationDir)

	dataSource := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, dbPort, username, dbName, password, sslMode)

	conn, err := sql.Open("postgres", dataSource)
	if err != nil {
		panic(err)
	}

	err = migrator.ApplyMigrations(conn)
	if err != nil {
		panic(err)
	}

	groupRepository := repository.NewGroupRepository(conn)
	studentRepository := repository.NewStudentRepository(conn)

	groupService := service.NewGroupService(groupRepository)
	studentService := service.NewStudentService(studentRepository)

	handler := handler.NewHandler(
		*groupService,
		*studentService,
	)

	srv := new(pkg.Server)

	if err := srv.Run(serverPort, handler.InitRoutes()); err != nil {
		log.Fatalf("can not run http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("application")

	return viper.ReadInConfig()
}
