package tasks

import (
	"testing"
	"time"

	"github.com/henrique502/golang-clean/domain/team"
	mocks "github.com/henrique502/golang-clean/mocks/domain/team"
	"github.com/stretchr/testify/assert"
)

func TestTaskTeam_SyncTeams(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		assertions := assert.New(t)

		nextURL := (*string)(nil)
		query := map[string]string{}

		data := team.Team{
			ID:          "1df96c22-faac-4151-af7e-a70a329f5e9a",
			Name:        "Ghostbusters",
			Description: "",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		list := []team.Team{}
		list = append(list, data)
		list = append(list, data)
		list = append(list, data)
		list = append(list, data)

		service := new(mocks.Service)
		service.On("GetTeamList", nextURL, query).Return(list, nil, nil)

		repository := new(mocks.Repository)
		repository.On("TeamUpSert", data).Return(nil)

		task := NewTaskTeam(repository, service)
		err := task.SyncTeams()

		assertions.Nil(err)
		service.AssertExpectations(t)
		service.AssertNumberOfCalls(t, "GetTeamList", 1)
		repository.AssertExpectations(t)
		repository.AssertNumberOfCalls(t, "TeamUpSert", 4)
	})
}
