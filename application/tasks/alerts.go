package tasks

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

func (s TaskAlertContainer) SyncAlerts() error {
	log.Debugln("Start teams sync")

	day := time.Now().AddDate(0, 0, -1)
	start := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, day.Location())
	end := time.Date(day.Year(), day.Month(), day.Day(), 23, 59, 59, 0, day.Location())

	nextURL := (*string)(nil)
	query := map[string]string{
		"order": "desc",
		"sort":  "createdAt",
		"query": fmt.Sprintf("createdAt > %d AND createdAt < %d", start.Unix(), end.Unix()),
	}

	for {
		list, next, err := s.service.GetAlertList(nextURL, query)
		if err != nil {
			return err
		}

		for _, element := range list {
			log.Infof("Fetch team %s", element.ID)
			err := s.repository.AlertUpSert(element)
			if err != nil {
				return err
			}
		}

		if next == nil {
			break
		}

		nextURL = next
	}

	return nil
}

func (s TaskAlertContainer) SyncAlertsAll() error {
	log.Debugln("Start teams sync")

	nextURL := (*string)(nil)
	query := map[string]string{
		"order": "desc",
		"sort":  "createdAt",
	}

	for {
		list, next, err := s.service.GetAlertList(nextURL, query)
		if err != nil {
			return err
		}

		for _, element := range list {
			log.Infof("Fetch team %s", element.ID)
			err := s.repository.AlertUpSert(element)
			if err != nil {
				return err
			}
		}

		if next == nil {
			break
		}

		nextURL = next
	}

	return nil
}
