package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq" // import your used driver
)
var ORM orm.Ormer
func ConnectToDb(){
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=birlcvav password=19QZsLOLiARafPCZOatrZobWlRYF5XGL dbname=birlcvav host=qdjjtnkv.db.elephantsql.com sslmode=disable")
	orm.RegisterModel(new(Users))
	ORM = orm.NewOrm()
}

func GetOrmObject() (o orm.Ormer) {
	return ORM
}
