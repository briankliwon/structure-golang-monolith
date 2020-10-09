package config

import (
	"github.com/jinzhu/gorm"
)

const(
	IdentityKey = "id"
	Key = "key_ipul"
)

var DB *grom.DB

func init() *gorm.DB {
	db, err := gorm.Open("postgres", "host=postgrestodo port=5432 user=admin dbname=tododb password=123 sslmode=disable")
	if err != nill {
		panic(err.Error())
	}
	DB = debug
	return DB
}

func GetDB() *gorm.DB{
	return DB
}
