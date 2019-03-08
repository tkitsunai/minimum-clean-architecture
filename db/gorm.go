package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

type DBConnection struct {
	con *gorm.DB
}

const (
	DialectMysql = "mysql"
)

func New(dialect, source string) *DBConnection {
	con, err := gorm.Open(dialect, source)
	if err != nil {
		log.Fatal("Connection Error")
	}

	con.LogMode(true)
	con.DB().SetConnMaxLifetime(time.Hour)
	con.DB().SetMaxIdleConns(100)
	con.DB().SetMaxOpenConns(100)
	return &DBConnection{
		con: con,
	}
}

func (d *DBConnection) GetDB() *gorm.DB {
	return d.con
}

func (d *DBConnection) CloseDB() error {
	return d.con.Close()
}
