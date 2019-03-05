package driver

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tkitsunai/minimum-clean-architecture/config"
)

type UserDriver struct{}

func NewUserDriver() *UserDriver {
	return &UserDriver{}
}

type UserModels []UserModel

type UserModel struct {
	Id   uint   `gorm:"primary_key"`
	Name string `gorm:"column:name"`
}

func (UserModel) TableName() string {
	return "user"
}

const (
	dialectMysql = "mysql"
)

func (u *UserDriver) FindAll() ([]UserModel, error) {
	db, err := gorm.Open(dialectMysql, config.DBConnection().Source)

	if err != nil {
		return nil, err
	}

	users := []UserModel{}

	if res := db.Find(&users); res.Error != nil {
		return nil, res.Error
	}

	return users, nil
}

func (u *UserDriver) SaveAll(users UserModels) error {
	db, err := gorm.Open(dialectMysql, config.DBConnection().Source)

	if err != nil {
		return err
	}

	db.Begin()

	var errors []error
	for _, user := range users {
		result := db.Create(user)
		if result.Error != nil {
			errors = append(errors, result.Error)
		}
	}

	if len(errors) > 0 {
		db.Rollback()
	}

	db.Commit()
	return nil
}

func (u *UserDriver) Save(user UserModel) error {
	return nil
}
