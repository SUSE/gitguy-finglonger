package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/SUSE/gitguy-finglonger/pkg/github/model"
)

// LabelActions moves an issue in the project to a columns depending on the label that has been set.
func (a *API) LabelActions(issue *model.Issue, w http.ResponseWriter) {
	switch issue.Label.Name {
	case "Blocked", "needinfo":
		possibleColumns := []int{a.Config.InProgressColumnID}
		notes, err := getNotesByColumns(a.Config.Github.APIURL, possibleColumns)
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
		status, _ := request("POST", url, notePayload)
		w.WriteHeader(status)
		return
	case "BugSquad":
		url := fmt.Sprintf(`%s/projects/columns/%d/cards`, a.Config.Github.APIURL, a.Config.TriagedColumnID)
		notePayload := model.NotePayload{
			ContentID:   issue.Issue.ID,
			ContentType: "Issue",
		}
		status, _ := request("POST", url, notePayload)
		w.WriteHeader(status)
		return
	}
}

// UnlabelActions moves an issue in the project to a columns depending on the label that has been unset.
func (a *API) UnlabelActions(issue *model.Issue, w http.ResponseWriter) {
	switch issue.Label.Name {
	case "needinfo":
		possibleColumns := []int{a.Config.BlockedColumnID}
		notes, err := getNotesByColumns(a.Config.Github.APIURL, possibleColumns)
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
		status, _ := request("POST", url, notePayload)
		w.WriteHeader(status)
		return
	}
}

// setLabels sets Labels to the issue
func (a *API) setLabels(issue *model.Issue, labels []string) error {
	url := fmt.Sprintf(`%s/repos/%s/%s/issues/%d/labels`, a.Config.Github.APIURL, issue.Repository.Owner.Login, issue.Repository.Name, issue.Issue.Number)
	labelsPayload := model.LabelsPayload{
		Labels: labels,
	}
	status, b := request("POST", url, labelsPayload)
	if status != 200 {
		return errors.New(string(b))
	}
	return nil
}

// removeLabel sets a Label to the issue
func (a *API) removeLabel(issue *model.Issue, label string) error {
	url := fmt.Sprintf(`%s/repos/%s/%s/issues/%d/labels/%s`, a.Config.Github.APIURL, issue.Repository.Owner.Login, issue.Repository.Name, issue.Issue.Number, label)
	status, b := request("DELETE", url, nil)
	if status != 200 {
		return errors.New(string(b))
	}
	return nil
}
