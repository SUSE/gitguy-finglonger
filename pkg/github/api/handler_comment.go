package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/SUSE/gitguy-finglonger/pkg/github/model"
	"github.com/SUSE/gitguy-finglonger/pkg/security"
)

// CommentHandler handles webhooks events from comments in github issues/PR
func (a *API) CommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	signature := r.Header.Get("X-Hub-Signature")
	if !security.IsValidSignature(body, signature) {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	comment := new(model.Comment)
	json.Unmarshal(body, comment)
}
