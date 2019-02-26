package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"path"
	"strconv"

	"github.com/chentex/github-project-mgr/pkg/github/model"
)

func getNotesByColumns(api string, columns []int) ([]model.NoteInColumns, error) {
	notes := []model.NoteInColumns{}
	for _, c := range columns {
		url := fmt.Sprintf(`%s/projects/columns/%d/cards`, api, c)
		status, bd := request("GET", url, nil)
		if status != 200 {
			return nil, errors.New("failed request")
		}
		var n []model.NoteInColumns
		err := json.Unmarshal(bd, &n)
		if err != nil {
			return nil, err
		}
		notes = append(notes, n...)
	}
	return notes, nil
}

func getCardID(notes []model.NoteInColumns, issueNumber int) (int, error) {
	cardID := 0
	for _, n := range notes {
		if n.ContentURL == "" {
			continue
		}
		url, err := url.Parse(n.ContentURL)
		if err != nil {
			return 0, err
		}
		p := url.Path
		in := path.Base(p)
		num, err := strconv.ParseInt(in, 10, 64)
		if err != nil {
			log.Fatalf("ParseInt: %s", err)
			return 0, err
		}
		if num == int64(issueNumber) {
			cardID = n.ID
			log.Printf("Card to move: %d", cardID)
			break
		}
	}
	if cardID == 0 {
		return 0, errors.New("No card found")
	}
	return cardID, nil
}
