package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/SUSE/gitguy-finglonger/pkg/github/model"
	"github.com/SUSE/gitguy-finglonger/pkg/security"
)

// IssueHandler handles webhooks events from issues in github
func (a *API) IssueHandler(w http.ResponseWriter, r *http.Request) {
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
	if !security.IsValidSignature(body, signature, a.Config.Github.Secret) {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	issue := new(model.Issue)
	json.Unmarshal(body, issue)
	log.Printf("process action: %s", issue.Action)

	switch issue.Action {
	case "labeled":
		a.LabelActions(issue, w)
	case "unlabeled":
		a.UnlabelActions(issue, w)
	case "assigned":
		a.MoveIssueInProgress(issue, w)
	case "unassigned":
		a.MoveIssueBlocked(issue, w)
	case "closed":
		a.ClosedIssue(issue, w)
	}
}
