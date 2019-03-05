package user

import "github.com/tkitsunai/minimum-clean-architecture/domain"

type UserPort interface {
	FindAllUser() (domain.Users, error)
}
