package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/SUSE/gitguy-finglonger/pkg/github/model"
)

// Review post a comment in an issue or pull request
func (a *API) Review(pr *model.PullRequest, w http.ResponseWriter) {
	url := fmt.Sprintf(`%s/repos/%s/%s/pulls/%d/reviews`, a.Config.Github.APIURL, pr.Repository.Owner.Login, pr.Repository.Name, pr.Number)
	commentPayload := model.ReviewCommentPayload{
		Body:  "## Release Coordination Checklist \r\n\r\n- [ ] item 1\r\n- [ ] item 2\r\n- [ ] item 3",
		Event: "REQUEST_CHANGES",
	}
	status, _ := request("POST", url, commentPayload)
	w.WriteHeader(status)
}

// SubmitReview approve a Review when check box are all completed
func (a *API) SubmitReview(r *model.PullRequestReview, w http.ResponseWriter) {
	url := fmt.Sprintf(`%s/repos/%s/%s/pulls/%d/reviews/%d/events`, a.Config.Github.APIURL, r.Repository.Owner.Login, r.Repository.Name, r.PullRequest.Number, r.Review.ID)
	if checkCheckList(r.Review.Body) {
		commentPayload := model.SubmitReviewPayload{
			Body:  "LGTM",
			Event: "APPROVE",
		}
		status, _ := request("POST", url, commentPayload)
		w.WriteHeader(status)
	}
}

func checkCheckList(s string) bool {
	items := strings.Count(s, "- [")
	completed := strings.Count(s, "- [x]")
	return items == completed
}
