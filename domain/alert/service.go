package alert

type Service interface {
	GetAlertList(nextURL *string, query map[string]string) ([]Alert, *string, error)
}
