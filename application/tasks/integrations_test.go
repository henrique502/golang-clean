package tasks

import (
	"testing"
	"time"

	"github.com/henrique502/golang-clean/domain/integration"
	mocks "github.com/henrique502/golang-clean/mocks/domain/integration"
	"github.com/stretchr/testify/assert"
)

func TestTaskIntegration_SyncIntegrations(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		assertions := assert.New(t)

		nextURL := (*string)(nil)
		query := map[string]string{}

		data := integration.Integration{
			ID:        "1df96c22-faac-4151-af7e-a70a329f5e9a",
			Name:      "API Git",
			Enabled:   true,
			Type:      "API",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		list := []integration.Integration{}
		list = append(list, data)
		list = append(list, data)
		list = append(list, data)
		list = append(list, data)

		service := new(mocks.Service)
		service.On("GetIntegrationList", nextURL, query).Return(list, nil, nil)

		repository := new(mocks.Repository)
		repository.On("IntegrationUpSert", data).Return(nil)

		task := NewTaskIntegration(repository, service)
		err := task.SyncIntegrations()

		assertions.Nil(err)
		service.AssertExpectations(t)
		service.AssertNumberOfCalls(t, "GetIntegrationList", 1)
		repository.AssertExpectations(t)
		repository.AssertNumberOfCalls(t, "IntegrationUpSert", 4)
	})
}
