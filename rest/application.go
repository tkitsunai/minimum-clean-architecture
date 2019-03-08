package rest

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tkitsunai/minimum-clean-architecture/config"
	"github.com/tkitsunai/minimum-clean-architecture/rest/user"
)

type Application struct {
	echo *echo.Echo
}

func (r *Application) Run() error {
	return r.echo.Start(config.AppPort())
}

func NewApplication(
	con *gorm.DB,
) *Application {

	e := echo.New()
	e.Add(rest.FactoryUserResource(con).GetUsers())
	e.Add(rest.FactoryUserResource(con).GetUserById())
	e.Use(middleware.Logger())

	return &Application{
		echo: e,
	}
}
