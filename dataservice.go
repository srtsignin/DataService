package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/srtsignin/dataService/router"
)

// PORT to expose for the service
const PORT = 3000

func main() {
	log.SetFlags(log.LUTC | log.Llongfile | log.Ldate | log.Ltime)
	log.Printf("Starting on port %d\n", PORT)
	router := router.NewRouter()

	log.Fatalln(http.ListenAndServe(":"+strconv.Itoa(PORT), router))
}
