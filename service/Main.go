package main

import (
	"log"
	"net/http"
	"strconv"
)

// PORT to expose for the service
const PORT = 3000

func main() {
	log.SetFlags(log.LUTC | log.Llongfile | log.Ldate | log.Ltime)
	log.Printf("Starting on port %d", PORT)
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), router))
}
