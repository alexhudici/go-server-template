package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
)

func init() {
	router.GET("/hello/:name", Hello)
	router.POST("/kong/:plugin", Kong)
}

// Hello func
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if log.Level == logrus.DebugLevel {
		dump, _ := httputil.DumpRequest(r, true)
		log.Debugf("%q", dump)
	}

	fmt.Fprintf(w, "hello, "+ps.ByName("name")) //%s!\n", ps.ByName("name"))
}

// Kong receiver func
func Kong(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Plugin, "+ps.ByName("plugin"))
	fmt.Println("Plugin, " + ps.ByName("plugin"))
}
