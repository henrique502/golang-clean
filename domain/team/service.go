package team

type Service interface {
	GetTeamList(nextURL *string, query map[string]string) ([]Team, *string, error)
}
