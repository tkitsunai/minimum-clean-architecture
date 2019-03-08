package driver

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserDriver struct {
	con *gorm.DB
}

func NewUserDriver(
	con *gorm.DB,
) *UserDriver {
	return &UserDriver{
		con: con,
	}
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
	users := []UserModel{}

	if res := u.con.Find(&users); res.Error != nil {
		return nil, res.Error
	}

	return users, nil
}

func (u *UserDriver) SaveAll(users UserModels) error {
	tx := u.con.Begin()

	if tx.Error != nil {
		return tx.Error
	}

	var errors []error
	for _, user := range users {
		result := tx.Create(user)
		if result.Error != nil {
			errors = append(errors, result.Error)
		}
	}

	defer rollBackHasErrors(errors, tx)
	defer tx.Commit()

	return nil
}

func rollBackHasErrors(errors []error, tx *gorm.DB) error {
	if errors == nil || len(errors) == 0 {
		return nil
	}
	return tx.Rollback().Error
}

func (u *UserDriver) Save(user UserModel) error {
	return nil
}
