package api

import (
	"fmt"
	"net/http"

	"github.com/SUSE/gitguy-finglonger/pkg/github/model"
)

// ClosedIssue when an issue is closed it should move to the Release Squad Board
func (a *API) ClosedIssue(issue *model.Issue, w http.ResponseWriter) {
	for _, l := range issue.Issue.Labels {
		if l.Name == "wontfix" {
			return
		}
	}

	url := fmt.Sprintf(`%s/projects/columns/%d/cards`, a.Config.Github.APIURL, a.Config.BacklogReleaseSquadColumnID)
	notePayload := model.NotePayload{
		ContentID:   issue.Issue.ID,
		ContentType: "Issue",
	}
	status, _ := request("POST", url, notePayload)
	w.WriteHeader(status)
}
