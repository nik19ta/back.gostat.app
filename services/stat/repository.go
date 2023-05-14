package stat

import "gostat/services/links"

type UserRepository interface {
	AddVisit(data Visits) error
	VisitExtend(session string) error
	GetVisits() ([]Visits, error)

	GetLinks() ([]links.Link, error)
}
