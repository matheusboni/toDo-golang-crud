package config

import (
	"fmt"
	"gorm.io/gorm/schema"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"toDo-golang-crud/model"
)

var DB *gorm.DB

func SetUpDB(env *Env) {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
			env.DataBaseHost,
			env.DataBaseUser,
			env.DataBasePassword,
			env.DataBaseName,
			env.DataBasePort),
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: env.DataBaseSchema + ".",
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	_ = db.AutoMigrate(model.ToDo{})
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}