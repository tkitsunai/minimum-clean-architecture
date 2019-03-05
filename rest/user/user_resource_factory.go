package rest

import (
	"github.com/tkitsunai/minimum-clean-architecture/driver"
	userGateway "github.com/tkitsunai/minimum-clean-architecture/gateway/user"
	presenter "github.com/tkitsunai/minimum-clean-architecture/presenter/user"
	"github.com/tkitsunai/minimum-clean-architecture/usecase/user"
)

func FactoryUserResource() *UserResource {
	uc, _ := user.NewUserUsecase(
		userGateway.NewUserGateway(
			driver.NewUserDriver(),
		),
		presenter.NewUserResponsePresenter(),
	)
	return newUserResource(uc)
}
