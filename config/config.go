package config

import (
// 	"fmt"
	"os"
	mgo "gopkg.in/mgo.v2"
  "log"
  "github.com/BurntSushi/toml"
)

var db *mgo.Database

type Config struct {
	Server   string
	Database string
}

func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}

func GetMongoDB() (*mgo.Database, error) {
	os.Setenv("MONGO_HOST", "mongodb://localhost/")
	os.Setenv("MONGO_DB_NAME", "go-rest-mongodb")

	host := os.Getenv("MONGO_HOST")
	dbName := os.Getenv("MONGO_DB_NAME")

	session, err := mgo.Dial(host)

	if err != nil {
		return nil, err
	}

	db := session.DB(dbName)

	return db, nil
}
