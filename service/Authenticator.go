package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/srtsignin/data-service/models"
	"github.com/tkanos/gonfig"
)

// Authenticate wraps HTTP requests with Authentication
func Authenticate(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokens := strings.Split(name, "_")
		if tokens[0] == "" {
			inner.ServeHTTP(w, r)
		} else {
			checkPermissions(w, r)
			inner.ServeHTTP(w, r)
		}
	})
}

func checkPermissions(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("AuthToken")
	if token == "" {
		log.Panicln("No Auth Token Provided")
	}

}

func callRoleService(token string) {
	conf := models.Configuration{}
	err := gonfig.GetConf("./config/data-service-conf.json", &conf)
	if err != nil {
		log.Fatalln("Misconfigured config")
	}

	req, err := http.NewRequest("GET", conf.RoleService.URL, nil)
	req.Header.Set("AuthToken", token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func auth(role []string, minimum string) bool {
	// for thing in roles if any of them bigger than minimum, good
	return getAuthValue(minimum) > 1
}

func getAuthValue(role string) int {
	if role == "Admin" {
		return 4
	}
	if role == "Staff" {
		return 3
	}
	if role == "Tutor" {
		return 2
	}
	return 1
}
