package convert

import (
	"github.com/tkitsunai/minimum-clean-architecture/domain"
	"github.com/tkitsunai/minimum-clean-architecture/driver"
)

func ToUsers(models driver.UserModels) domain.Users {
	var users domain.Users
	for _, model := range models {
		users = append(users, domain.User{
			Name: model.Name,
		})
	}
	return users
}
