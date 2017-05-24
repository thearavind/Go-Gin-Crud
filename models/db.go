package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq" //Postgres driver
)

var ormObject orm.Ormer

// ConnectToDb - Initializes the ORM and Connection to the postgres DB
func ConnectToDb() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=birlcvav password=19QZsLOLiARafPCZOatrZobWlRYF5XGL dbname=birlcvav host=qdjjtnkv.db.elephantsql.com sslmode=disable")
	orm.RegisterModel(new(Users))
	ormObject = orm.NewOrm()
}

// GetOrmObject - Getter function for the ORM object with which we can query the database
func GetOrmObject() (o orm.Ormer) {
	return ormObject
}
