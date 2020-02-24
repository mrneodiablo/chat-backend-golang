package db

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

var Path, Name string

func init() {
	Path = beego.AppConfig.String("DBPath")
	Name = beego.AppConfig.String("DBName")
}

type Database struct {
	s       *mgo.Session
	name    string
	session *mgo.Database
}

func (db *Database) Connect() {
	db.s = service.Session()
	session := *db.s.DB(db.name)
	db.session = &session
}

func DropDatabase() {
	session, err := mgo.Dial(Path)
	defer session.Close()
	if err != nil {
		panic(err)
	}

	err = session.DB(Name).DropDatabase()
	if err != nil {
		panic(err)
	}
}

func newDBSession(name string) *Database {
	var db = Database{
		name: name,
	}
	db.Connect()
	return &db
}