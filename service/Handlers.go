package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/srtsignin/DataService/database"
	"github.com/srtsignin/DataService/models"
)

// Index provides the heartbeat endpoint to determine availability of service
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

// Store handles POST requests to /store
// This converts the object to a checkoff model and
// sends it to long term storage
func Store(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if recovery := recover(); recovery != nil {
			fmt.Fprintln(w, models.CreateHTTPResponse(recovery, false).ToJSON())
		}
	}()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panicln(err)
	}

	log.Printf("Received in body: %v\n", string(body))
	activeUserModel := models.ActiveUserModel{}
	json.Unmarshal(body, &activeUserModel)

	db := database.GetDriver()
	db.Store(models.CreateCheckoff(activeUserModel))
}
