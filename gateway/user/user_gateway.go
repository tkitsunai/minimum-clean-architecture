package user

import (
	"github.com/tkitsunai/minimum-clean-architecture/domain"
	"github.com/tkitsunai/minimum-clean-architecture/driver"
	"github.com/tkitsunai/minimum-clean-architecture/gateway/user/convert"
)

// port実装
type UserGateway struct {
	driver *driver.UserDriver
}

func NewUserGateway(
	driver *driver.UserDriver,
) *UserGateway {
	return &UserGateway{
		driver: driver,
	}
}

func (u *UserGateway) FindAllUser() (domain.Users, error) {
	models, err := u.driver.FindAll()

	if err != nil {
		return domain.NewEmptyUsers(), err
	}

	return convert.ToUsers(models), nil
}
