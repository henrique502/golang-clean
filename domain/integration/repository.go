package integration

type Repository interface {
	IntegrationUpSert(integration Integration) error
}
