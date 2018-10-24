package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/srtsignin/data-service/service"
)

// PORT to expose for the service
const PORT = 3000

func main() {
	log.SetFlags(log.LUTC | log.Llongfile | log.Ldate | log.Ltime)
	log.Printf("Starting on port %d\n", PORT)
	router := service.NewRouter()

	log.Fatalln(http.ListenAndServe(":"+strconv.Itoa(PORT), router))
}
