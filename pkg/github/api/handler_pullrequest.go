package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/SUSE/gitguy-finglonger/pkg/github/model"
	"github.com/SUSE/gitguy-finglonger/pkg/security"
)

// PullRequestHandler handles webhooks events from comments in github issues/PR
func (a *API) PullRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("%v", r.Header)
	signature := r.Header.Get("X-Hub-Signature")
	if !security.IsValidSignature(body, signature) {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	event := r.Header.Get("X-Github-Event")
	switch event {
	case "pull_request":
		pullRequest := new(model.PullRequest)
		json.Unmarshal(body, pullRequest)
		switch pullRequest.Action {
		case "opened":
			a.Review(pullRequest, w)
		}
	case "pull_request_review":
		review := new(model.PullRequestReview)
		json.Unmarshal(body, review)
		switch review.Action {
		case "edited":
			a.SubmitReview(review, w)
		}
	}
}
