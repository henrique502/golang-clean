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
	service := opsgenie.New(nil)
	err := tasks.NewTaskTeam(repository, service).SyncTeams()

	if err != nil {
		log.Fatalln(err)
	} else {
		os.Exit(0)
	}
}
