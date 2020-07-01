package api

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/seanarwa/common/config"
	"github.com/seanarwa/ds/api/health"
	"github.com/seanarwa/ds/api/person"
	"github.com/seanarwa/ds/util"
	log "github.com/sirupsen/logrus"
)

func Start() {

	port := config.GetString("api.port")
	logFields := log.Fields{"port": port}

	setAllHandlers()
	log.WithFields(logFields).Debug("HTTP server has started")

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.WithFields(logFields).Fatal(
			"Error occured when trying to start HTTP server: ",
			http.ListenAndServe(":"+port, nil),
		)
	}()

	<-done
	log.WithFields(logFields).Debug("HTTP server has stopped")
}

func setAllHandlers() {
	http.HandleFunc("/", allHandler)
	http.HandleFunc("/health", health.Handler)
	http.HandleFunc("/person", person.Handler)
}

func allHandler(w http.ResponseWriter, req *http.Request) {

	method := req.Method
	path := req.URL.EscapedPath()

	log.Debug(method, " ", path)

	if path == "/" {
		rootHandler(w, req)
	} else {
		errorNotFoundHandler(w, req)
	}
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		util.WriteHTTPResponse(w, map[string]interface{}{
			"name":    config.GetString("name"),
			"version": config.GetString("version"),
		}, http.StatusAccepted)
	}
}

func errorNotFoundHandler(w http.ResponseWriter, req *http.Request) {
	util.WriteHTTPResponse(w, map[string]interface{}{
		"message": "404 path not found " + req.URL.EscapedPath(),
	}, http.StatusNotFound)
}
