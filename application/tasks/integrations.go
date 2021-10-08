package tasks

import (
	log "github.com/sirupsen/logrus"
)

func (s TaskIntegrationContainer) SyncIntegrations() error {
	log.Debugln("Start teams sync")

	nextURL := (*string)(nil)
	query := map[string]string{}

	for {
		list, next, err := s.service.GetIntegrationList(nextURL, query)
		if err != nil {
			return err
		}

		for _, element := range list {
			log.Infof("Fetch team %s", element.Name)
			err := s.repository.IntegrationUpSert(element)
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
