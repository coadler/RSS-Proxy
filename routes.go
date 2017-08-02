package main

import "net/http"

// Route is the structure for reach route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes stores all routes in a slice
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Retrieve RSS",
		"POST",
		"/v1/get",
		GetRSS,
	},
}
