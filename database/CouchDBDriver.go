package database

import (
	"log"
	"strings"
	"time"

	couchdb "github.com/rhinoman/couchdb-go"
	"github.com/srtsignin/data-service/models"
	"github.com/tkanos/gonfig"
)

// CouchDBDriver provides a LTS driver for interacting
// with CouchDB
type CouchDBDriver struct {
}

// Store stores a Checkoff model into CouchDB
func (couchDbDatabase CouchDBDriver) Store(checkoff models.Checkoff) {
	db := connectToCouchDB()

	_, err := db.Save(checkoff, checkoff.GetID(), "")

	if err != nil {
		log.Panicln(err)
	}
}

// GenerateCSV reads from the database and creates a .csv string
func (couchDbDatabase CouchDBDriver) GenerateCSV() string {
	db := connectToCouchDB()

	var results AllDocsResponse

	err := db.GetView("all_docs", "all_docs", &results, nil)

	if err != nil {
		log.Panicln(err)
	}
	log.Printf("Received %d results from couchdb", results.TotalRows)
	log.Printf("Returned from Couchdb: %+v\n", results)

	var builder strings.Builder

	for _, row := range results.Rows {
		builder.WriteString(row.Data.Delimited(','))
		builder.WriteRune('\n')
	}

	return builder.String()
}

func connectToCouchDB() *couchdb.Database {
	conf := models.Configuration{}
	err := gonfig.GetConf("./config/data-service-conf.json", &conf)
	if err != nil {
		log.Panicln(err)
	}

	var timeout = time.Duration(500 * time.Millisecond)
	conn, err := couchdb.NewConnection(conf.CouchDB.ConnectionString, conf.CouchDB.DatabasePort, timeout)
	auth := couchdb.BasicAuth{Username: conf.CouchDB.Username, Password: conf.CouchDB.Password}

	if err != nil {
		log.Panicln(err)
	}

	return conn.SelectDB(conf.CouchDB.Database, &auth)
}
