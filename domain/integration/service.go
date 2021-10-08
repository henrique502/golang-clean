package integration

type Service interface {
	GetIntegrationList(nextURL *string, query map[string]string) ([]Integration, *string, error)
}
