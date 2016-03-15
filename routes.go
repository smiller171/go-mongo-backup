package main

import "net/http"

//Route defines route struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes defines route var type
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"Get",
		"/",
		index,
	},
	Route{
		"DumpCreate",
		"POST",
		"/v0/dump",
		dumpCreate,
	},
}
