package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/henrique502/golang-clean/application/tasks"
	"github.com/henrique502/golang-clean/infrastructure/database"
	"github.com/henrique502/golang-clean/infrastructure/external/opsgenie"
	_ "github.com/henrique502/golang-clean/infrastructure/logger"
)

func main() {
	repository := database.New()
	defer repository.Close()

	service := opsgenie.New(nil)
	err := tasks.NewTaskIntegration(repository, service).SyncIntegrations()

	if err != nil {
		log.Fatalln(err)
	} else {
		os.Exit(0)
	}
}
