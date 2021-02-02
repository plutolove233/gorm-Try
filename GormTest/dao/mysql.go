package dao

import(
	"GormTest/configs"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Startsql()(err error){
	DB,err = gorm.Open(configs.GetConfigInf())
	return err
}

type DBmodel interface{
	Add() error
	Delete() error
	Updata() error
	Get() error
}