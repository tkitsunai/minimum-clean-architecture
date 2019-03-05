package user

import (
	"github.com/tkitsunai/minimum-clean-architecture/port/response"
	"github.com/tkitsunai/minimum-clean-architecture/port/user"
)

type UserUsecase struct {
	userPort     user.UserPort
	responsePort response.UserResponsePort
}

func NewUserUsecase(
	userPort user.UserPort,
	responsePort response.UserResponsePort,
) (*UserUsecase, error) {
	return &UserUsecase{
		userPort:     userPort,
		responsePort: responsePort,
	}, nil
}

func (u *UserUsecase) GetUsers(ctx response.JsonContext) error {
	if users, err := u.userPort.FindAllUser(); err != nil {
		return err
	} else {
		return u.responsePort.Response(users, ctx)
	}
}
