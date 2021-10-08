package alert

type Repository interface {
	AlertUpSert(alert Alert) error
}
