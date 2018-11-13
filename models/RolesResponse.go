package models

// RolesResponse is a struct for de-serializing the json response
// from the roles service
type RolesResponse struct {
	User    User     `json:"user"`
	Roles   []string `json:"roles"`
	Message string   `json:"message"`
}

// User is a wrapper for the username and name of a person
type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}
