package rest

import (
	"github.com/labstack/echo"
	"github.com/tkitsunai/minimum-clean-architecture/usecase/user"
	"net/http"
	"path"
)

type UserResource struct {
	basePath string
	usecase  *user.UserUsecase
}

func newUserResource(
	usecase *user.UserUsecase,
) *UserResource {
	return &UserResource{
		basePath: "v1/users",
		usecase:  usecase,
	}
}

func (u *UserResource) GetUsers() (method string, resourcePath string, handlerFunc echo.HandlerFunc) {
	return http.MethodGet,
		path.Join(u.basePath, ""),
		func(ctx echo.Context) error {
			return u.usecase.GetUsers(ctx)
		}
}

func (u *UserResource) GetUserById() (method string, resourcePath string, handlerFunc echo.HandlerFunc) {
	return http.MethodGet,
		path.Join(u.basePath, ":id"),
		func(ctx echo.Context) error {
			return ctx.JSON(http.StatusNotImplemented, "")
		}
}
