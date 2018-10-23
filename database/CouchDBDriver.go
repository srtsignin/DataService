package database

import (
	"DataService/models"
	"log"
	"time"

	couchdb "github.com/rhinoman/couchdb-go"
	"github.com/tkanos/gonfig"
)

// CouchDBDriver provides a LTS driver for interacting
// with CouchDB
type CouchDBDriver struct {
}

// Store stores a Checkoff model into CouchDB
func (couchDbDatabase CouchDBDriver) Store(checkoff models.Checkoff) {
	conf := models.Configuration{}
	err := gonfig.GetConf("./config/data-service-conf.json", &conf)
	if err != nil {
		log.Panicln(err)
	}

	var timeout = time.Duration(500 * time.Millisecond)
	conn, err := couchdb.NewConnection(conf.CouchDB.ConnectionString,
		conf.CouchDB.DatabasePort, timeout)
	if err != nil {
		log.Panicln(err)
	}

	auth := couchdb.BasicAuth{Username: conf.CouchDB.Username, Password: conf.CouchDB.Password}
	db := conn.SelectDB(conf.CouchDB.Database, &auth)

	_, err = db.Save(checkoff, checkoff.GetID(), "")

	if err != nil {
		log.Panicln(err)
	}
}
