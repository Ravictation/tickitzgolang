package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Ravictation/tickitzgolang/internal/pkg"
	"github.com/Ravictation/tickitzgolang/internal/routers"
	"github.com/asaskevich/govalidator"
	"github.com/spf13/viper"

	_ "github.com/joho/godotenv/autoload"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
	viper.SetConfigName("env.dev")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

}

func main() {
	serverFlag := flag.Bool("listen", false, "Run function a (server)")
	migrateUpFlag := flag.Bool("migrate-up", false, "Run function to apply migration")
	migrateDownFlag := flag.Bool("migrate-down", false, "Run function to rollback migration")

	flag.Parse()
	migrate := pkg.NewMigrator()

	if *serverFlag {
		listen()
	} else if *migrateUpFlag {
		migrate.Ups()
	} else if *migrateDownFlag {
		migrate.Downs()
	} else {
		fmt.Println("Usage: go run main.go --listen or go run main.go --migrate-up")
	}
}

func listen() {
	database, err := pkg.Pgdb()
	if err != nil {
		log.Fatal(err)
	}
	router := routers.New(database)
	server := pkg.Server(router)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

// migrate -path ./migrations -database "postgresql://tickitz:tickitzgolang@localhost/tickitz?port=5432&sslmode=disable&search_path=public" -verbose up
