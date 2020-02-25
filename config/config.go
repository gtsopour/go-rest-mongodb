package config

import (
	"github.com/BurntSushi/toml"
	mgo "gopkg.in/mgo.v2"
	"log"
)

type Config struct {
	Port	 string
	Server   string
	Database string
}

var config = Config{}
var db *mgo.Database

func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}

func GetDB() (*mgo.Database, error) {
	config.Read()

	host := config.Server
	dbName := config.Database

	session, err := mgo.Dial(host)

	if err != nil {
		return nil, err
	}

	db := session.DB(dbName)

	return db, nil
}
