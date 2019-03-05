package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func request(method string, url string, payload interface{}, token string) (int, []byte) {
	nb, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln(err)
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(nb))
	t := fmt.Sprintf(`token %s`, token)
	req.Header.Add("Accept", "application/vnd.github.inertia-preview+json")
	req.Header.Add("Authorization", t)
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("request error: %d - %s", resp.StatusCode, err.Error())
		log.Fatalln(err)
		return resp.StatusCode, nil
	}
	bd, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read respose Body: %s", err)
		return http.StatusInternalServerError, nil
	}
	log.Printf("request %s status: %d", url, resp.StatusCode)
	return resp.StatusCode, bd
}
