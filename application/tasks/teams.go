package tasks

import (
	log "github.com/sirupsen/logrus"
)

// Teams fetch and update teams table
func (s TaskTeamContainer) SyncTeams() error {
	log.Debugln("Start teams sync")

	nextURL := (*string)(nil)
	query := map[string]string{}

	for {
		list, next, err := s.service.GetTeamList(nextURL, query)
		if err != nil {
			return err
		}

		for _, element := range list {
			log.Infof("Fetch team %s", element.Name)
			err := s.repository.TeamUpSert(element)
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
