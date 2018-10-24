package service

import "net/http"

// A Route is a route exposed by the webservice
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is an array of multiple Route objects
type Routes []Route

var routes = Routes{
	Route{
		"Test",
		"GET",
		"/",
		Index,
	},
	Route{
		"Store",
		"POST",
		"/store",
		Store,
	},
}
