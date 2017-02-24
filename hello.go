package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
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
	fmt.Println("X-Auth-Token Header, " + r.Header.Get("X-Auth-Token"))
	fmt.Println("X-Auth-Token Header, " + r.Header.Get("X-Auth-Project-Id"))
	fmt.Println("X-Auth-Token Header, " + r.Header.Get("X-Auth-Service"))
}
