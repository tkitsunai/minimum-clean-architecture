package user

import (
	"github.com/tkitsunai/minimum-clean-architecture/domain"
	"github.com/tkitsunai/minimum-clean-architecture/port/response"
	"github.com/tkitsunai/minimum-clean-architecture/presenter/user/convert"
)

type UserResponsePresenter struct{}

func NewUserResponsePresenter() *UserResponsePresenter {
	return &UserResponsePresenter{}
}

func (u *UserResponsePresenter) Response(users domain.Users, jsonContext response.JsonContext) error {
	return jsonContext.JSON(200, convert.ToUserJSON(users))
}
