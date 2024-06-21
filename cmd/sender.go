//go:build sender

package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"

	"github.com/ShipIM/go-group-manager/api/rest/handler"
	"github.com/ShipIM/go-group-manager/internal/postgres"
	"github.com/ShipIM/go-group-manager/internal/rabbit"
	"github.com/ShipIM/go-group-manager/internal/repository"
	"github.com/ShipIM/go-group-manager/internal/service"
	"github.com/ShipIM/go-group-manager/pkg"

	"github.com/spf13/viper"
)

//go:embed migration/*.sql
var MigrationsFS embed.FS

func main() {
	if err := initSenderConfig(); err != nil {
		log.Fatalf("can not initialize configs: %s", err.Error())
	}

	var (
		serverPort      = viper.GetString("server.port")
		dbHost          = viper.GetString("db.host")
		dbPort          = viper.GetString("db.port")
		dbUsername      = viper.GetString("db.username")
		dbPassword      = viper.GetString("db.password")
		dbName          = viper.GetString("db.name")
		dbSslMode       = viper.GetString("db.sslmode")
		dbMigrationDir  = viper.GetString("db.migration_dir")
		rabbitAddress   = viper.GetString("rabbit.address")
		rabbitExchanger = viper.GetString("rabbit.exchanger")
	)

	migrator := postgres.NewMigrator(MigrationsFS, dbMigrationDir)

	dataSource := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbHost, dbPort, dbUsername, dbName, dbPassword, dbSslMode)

	conn_db, err := sql.Open("postgres", dataSource)
	if err != nil {
		panic(err)
	}

	err = migrator.ApplyMigrations(conn_db)
	if err != nil {
		panic(err)
	}

	conn_mq, err := rabbit.InitSenderRabbit(rabbitAddress, rabbitExchanger)
	if err != nil {
		panic(err)
	}
	defer conn_mq.Close()

	groupRepository := repository.NewGroupRepository(conn_db)
	studentRepository := repository.NewStudentRepository(conn_db)

	groupService := service.NewGroupService(groupRepository)
	studentService := service.NewStudentService(studentRepository, conn_mq, rabbitExchanger)

	handler := handler.NewHandler(
		*groupService,
		*studentService,
	)

	srv := new(pkg.Server)

	if err := srv.Run(serverPort, handler.InitRoutes()); err != nil {
		log.Fatalf("can not run http server: %s", err.Error())
	}
}

func initSenderConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("application-sender")

	return viper.ReadInConfig()
}
