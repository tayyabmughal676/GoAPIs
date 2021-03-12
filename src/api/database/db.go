package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"src/config"
)

func Connect() *gorm.DB  {
	//db, err := gorm.Open(mysql.New(config.DBDRIVER))
	db, err := gorm.Open(mysql.Open(config.DBURL), &gorm.Config{})



}
