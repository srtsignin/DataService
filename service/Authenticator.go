package service

import (
	"encoding/json"
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
		defer func() {
			if recovery := recover(); recovery != nil {
				log.Println(recovery)
				fmt.Fprintln(w, models.CreateHTTPResponse(recovery, false).ToJSON())
			}
		}()

		tokens := strings.Split(name, "_")
		if tokens[0] == "" {
			inner.ServeHTTP(w, r)
		} else {
			isAuthorized, response := checkPermissions(w, r, tokens[0])

			if !isAuthorized {
				log.Panicf("User %s is not authorized to access %s", response.User.Name, tokens[1])
			}
			inner.ServeHTTP(w, r)
		}
	})
}

func checkPermissions(w http.ResponseWriter, r *http.Request, minimumAccessLevel string) (bool, models.RolesResponse) {
	token := r.Header.Get("AuthToken")
	if token == "" {
		log.Panicln("No Auth Token Provided")
	}
	response := callRoleService(token)
	log.Printf("User received was %+v\n", response)
	return auth(response.Roles, minimumAccessLevel), response
}

func callRoleService(token string) models.RolesResponse {
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
	log.Println("response Body:", string(body))
	response := models.RolesResponse{}
	json.Unmarshal(body, &response)
	return response
}

func auth(roles []string, minimum string) bool {
	min := getAuthValue(minimum)
	for _, role := range roles {
		if getAuthValue(role) >= min {
			return true
		}
	}
	return false
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
