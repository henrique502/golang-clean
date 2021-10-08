package alert

import "time"

const TableName = "alerts"

type Alert struct {
	ID              string
	Priority        string
	Source          string
	Message         string
	ReportAckTime   int
	ReportCloseTime int
	IntegrationID   string
	ColletedAt      time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func New() Alert {
	return Alert{}
}
