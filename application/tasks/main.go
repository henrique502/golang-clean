package tasks

import (
	"github.com/henrique502/golang-clean/domain/alert"
	"github.com/henrique502/golang-clean/domain/integration"
	"github.com/henrique502/golang-clean/domain/team"
)

type TaskAlertContainer struct {
	repository alert.Repository
	service    alert.Service
}

type TaskIntegrationContainer struct {
	repository integration.Repository
	service    integration.Service
}

type TaskTeamContainer struct {
	repository team.Repository
	service    team.Service
}

type Task interface {
	SyncAlerts() error
	SyncAlertsAll() error
	SyncIntegrations() error
	SyncTeams() error
}

func NewTaskAlert(repository alert.Repository, service alert.Service) *TaskAlertContainer {
	return &TaskAlertContainer{
		repository: repository,
		service:    service,
	}
}

func NewTaskIntegration(repository integration.Repository, service integration.Service) *TaskIntegrationContainer {
	return &TaskIntegrationContainer{
		repository: repository,
		service:    service,
	}
}

func NewTaskTeam(repository team.Repository, service team.Service) *TaskTeamContainer {
	return &TaskTeamContainer{
		repository: repository,
		service:    service,
	}
}
