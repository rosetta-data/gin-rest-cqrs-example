package config

import (
	"github.com/kyhsa93/go-rest-example/study/model"
	"log"

	"github.com/caarlos0/env"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	Port     string `env:"DATABASE_PORT" envDefault:"3306"`
	Host     string `env:"DATABASE_HOST" envDefault:"localhost"`
	Name     string `env:"DATABASE_NAME" envDefault:"study"`
	User     string `env:"DATABASE_USER" envDefault:"root"`
	Password string `env:"DATABASE_PASSWORD" envDefault:"test"`
}

func GetConnection() *gorm.DB {
	database := Database{}
	env.Parse(&database)

	User := database.User
	Password := database.Password
	Name := database.Name
	Host := database.Host
	Port := database.Port

	db, err := gorm.Open("mysql", User+":"+Password+"@tcp("+Host+":"+Port+")/"+Name+"?parseTime=true")

	if err != nil {
		log.Println(err)
		panic(err)
	}

	db.AutoMigrate(model.Study{})

	db.LogMode(true)
	return db
}
