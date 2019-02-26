package api

import (
	"encoding/json"
	"log"
	"net/http"
)

var (
	version string
)

// VersionHandler returns version of the current api deployed
func (a *API) VersionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	v, err := json.Marshal(version)
	if err != nil {
		log.Fatalln(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	w.WriteHeader(200)
	w.Write(v)
}
