package team

type Repository interface {
	TeamUpSert(team Team) error
}
