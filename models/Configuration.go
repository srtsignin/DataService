package models

// Configuration is the overall configuration
type Configuration struct {
	CouchDB CouchDBConfiguration `json:"couchDB"`
}

// CouchDBConfiguration represents a configuration for
// couchDB
type CouchDBConfiguration struct {
	DatabasePort     int    `json:"databasePort"`
	ConnectionString string `json:"connectionString"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	Database         string `json:"database"`
}
