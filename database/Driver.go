package database

import (
	"DataService/models"
)

// Driver represents a driver object for connecting
// to a Long Term Storage database
type Driver interface {
	Store(checkoff models.Checkoff)
}

// GetDriver Returns a database driver
func GetDriver() Driver {
	return CouchDBDriver{}
}