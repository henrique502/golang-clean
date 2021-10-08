package tasks

import (
	"fmt"
	"testing"
	"time"

	"github.com/henrique502/golang-clean/domain/alert"
	mocks "github.com/henrique502/golang-clean/mocks/domain/alert"
	"github.com/stretchr/testify/assert"
)

func TestTaskAlert_SyncAlerts(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		assertions := assert.New(t)

		day := time.Now().AddDate(0, 0, -1)
		start := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, day.Location())
		end := time.Date(day.Year(), day.Month(), day.Day(), 23, 59, 59, 0, day.Location())

		nextURL := (*string)(nil)
		query := map[string]string{
			"order": "desc",
			"sort":  "createdAt",
			"query": fmt.Sprintf("createdAt > %d AND createdAt < %d", start.Unix(), end.Unix()),
		}

		data := alert.Alert{
			ID:              "cca0297f-0505-45c2-a5f2-bc6ee4e393f8-1628824646683",
			Priority:        "P2",
			Source:          "CloudWatch",
			Message:         "[P2] [Triggered] Erro ao gerar UR do recebedor",
			ReportAckTime:   43304,
			ReportCloseTime: 342864,
			IntegrationID:   "594cdb21-a9fa-4319-a010-6cdae3a473e1",
			ColletedAt:      time.Now(),
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}

		list := []alert.Alert{}
		list = append(list, data)
		list = append(list, data)
		list = append(list, data)
		list = append(list, data)

		service := new(mocks.Service)
		service.On("GetAlertList", nextURL, query).Return(list, nil, nil)

		repository := new(mocks.Repository)
		repository.On("AlertUpSert", data).Return(nil)

		task := NewTaskAlert(repository, service)
		err := task.SyncAlerts()

		assertions.Nil(err)
		service.AssertExpectations(t)
		service.AssertNumberOfCalls(t, "GetAlertList", 1)
		repository.AssertExpectations(t)
		repository.AssertNumberOfCalls(t, "AlertUpSert", 4)
	})
}

func TestTaskAlert_SyncAlertsAll(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		assertions := assert.New(t)

		nextURL := (*string)(nil)
		query := map[string]string{
			"order": "desc",
			"sort":  "createdAt",
		}

		data := alert.Alert{
			ID:              "cca0297f-0505-45c2-a5f2-bc6ee4e393f8-1628824646683",
			Priority:        "P2",
			Source:          "CloudWatch",
			Message:         "[P2] [Triggered] Erro ao gerar UR do recebedor",
			ReportAckTime:   43304,
			ReportCloseTime: 342864,
			IntegrationID:   "594cdb21-a9fa-4319-a010-6cdae3a473e1",
			ColletedAt:      time.Now(),
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}

		list := []alert.Alert{}
		list = append(list, data)
		list = append(list, data)
		list = append(list, data)
		list = append(list, data)

		service := new(mocks.Service)
		service.On("GetAlertList", nextURL, query).Return(list, nil, nil)

		repository := new(mocks.Repository)
		repository.On("AlertUpSert", data).Return(nil)

		task := NewTaskAlert(repository, service)
		err := task.SyncAlertsAll()

		assertions.Nil(err)
		service.AssertExpectations(t)
		service.AssertNumberOfCalls(t, "GetAlertList", 1)
		repository.AssertExpectations(t)
		repository.AssertNumberOfCalls(t, "AlertUpSert", 4)
	})
}

func TestTaskAlert_SyncAlertsAllPaginate(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		assertions := assert.New(t)

		nextURL := (*string)(nil)
		page2 := "http://localhost/p2"
		page3 := "http://localhost/p3"
		query := map[string]string{
			"order": "desc",
			"sort":  "createdAt",
		}

		data := alert.Alert{
			ID:              "cca0297f-0505-45c2-a5f2-bc6ee4e393f8-1628824646683",
			Priority:        "P2",
			Source:          "CloudWatch",
			Message:         "[P2] [Triggered] Erro ao gerar UR do recebedor",
			ReportAckTime:   43304,
			ReportCloseTime: 342864,
			IntegrationID:   "594cdb21-a9fa-4319-a010-6cdae3a473e1",
			ColletedAt:      time.Now(),
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}

		list := []alert.Alert{}
		list = append(list, data)
		list = append(list, data)
		list = append(list, data)
		list = append(list, data)

		service := new(mocks.Service)
		service.On("GetAlertList", nextURL, query).Return(list, &page2, nil)
		service.On("GetAlertList", &page2, query).Return(list, &page3, nil)
		service.On("GetAlertList", &page3, query).Return(list, nil, nil)

		repository := new(mocks.Repository)
		repository.On("AlertUpSert", data).Return(nil)

		task := NewTaskAlert(repository, service)
		err := task.SyncAlertsAll()

		assertions.Nil(err)
		service.AssertExpectations(t)
		service.AssertNumberOfCalls(t, "GetAlertList", 3)
		repository.AssertExpectations(t)
		repository.AssertNumberOfCalls(t, "AlertUpSert", 12)
	})
}
