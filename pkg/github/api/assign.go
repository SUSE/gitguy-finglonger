package api

import (
	"fmt"
	"net/http"

	"github.com/SUSE/gitguy-finglonger/pkg/github/model"
)

// MoveIssueInProgress when an issue has been assigned to someone
func (a *API) MoveIssueInProgress(issue *model.Issue, w http.ResponseWriter) {
	switch len(issue.Issue.Assignees) > 1 {
	case false:
		possibleColumns := []int{a.Config.TriagedColumnID, a.Config.BlockedColumnID}
		notes, err := a.getNotesByColumns(a.Config.Github.APIURL, possibleColumns)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		cardID, err := getCardID(notes, issue.Issue.Number)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		url := fmt.Sprintf(`%s/projects/columns/cards/%d/moves`, a.Config.Github.APIURL, cardID)
		notePayload := model.MoveNotePayload{
			Position: "top",
			ColumnID: a.Config.InProgressColumnID,
		}
		status, _ := request("POST", url, notePayload, a.Config.Github.Token)
		w.WriteHeader(status)
	case true:
		err := a.checkLabels(issue)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}
}

func (a *API) checkLabels(issue *model.Issue) error {
	found := false
	for _, l := range issue.Issue.Labels {
		if l.Name == "help wanted" {
			found = true
		}
	}
	if found && len(issue.Issue.Assignees) > 1 {
		err := a.setLabels(issue, []string{"pairing"})
		if err != nil {
			return err
		}
		err = a.removeLabel(issue, "help wanted")
		if err != nil {
			return err
		}
	}
	return nil
}
