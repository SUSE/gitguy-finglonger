package api

import (
	"fmt"
	"net/http"

	"github.com/SUSE/gitguy-finglonger/pkg/github/model"
)

// MoveIssueBlocked when an issue has been unassigned completely
func (a *API) MoveIssueBlocked(issue *model.Issue, w http.ResponseWriter) {
	if len(issue.Issue.Assignees) != 0 {
		return
	}
	possibleColumns := []int{a.Config.InProgressColumnID}
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
		ColumnID: a.Config.BlockedColumnID,
	}
	status, _ := request("POST", url, notePayload, a.Config.Github.Token)
	w.WriteHeader(status)
	return
}
