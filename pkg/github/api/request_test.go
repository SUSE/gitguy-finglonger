package api

import (
	"context"
	"log"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func Test_request(t *testing.T) {
	svr := startHTTPServer()
	waitForServer(svr.Addr)
	type args struct {
		method  string
		url     string
		payload interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []byte
	}{
		{"1", args{"POST", "http://127.0.0.1:9898/post", nil}, 200, []byte("")},
		{"2", args{"GET", "http://127.0.0.1:9898/get", nil}, 200, []byte("hello")},
		{"3", args{"POST", "http://127.0.0.1:9898/asd", nil}, 404, []byte("")},
		{"4", args{"POST", "http://127.0.0.1:9898/asd", nil}, 404, []byte("")},
		{"5", args{"GET", "http://127.0.0.1:9898/post", nil}, 404, []byte("")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := request(tt.args.method, tt.args.url, tt.args.payload)
			status := got != tt.want
			if status {
				t.Errorf("request() got = %v, want %v", got, tt.want)
			}
			if status && !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("request() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
	if err := svr.Shutdown(context.Background()); err != nil {
		t.Errorf("error stoping down web server")
	}
}

func startHTTPServer() *http.Server {
	srv := &http.Server{Addr: ":9898"}
	handler := http.NewServeMux()
	handler.HandleFunc("/post", post)
	handler.HandleFunc("/get", get)
	srv.Handler = handler
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()
	return srv
}

func waitForServer(url string) {
	var retry = 10
	for i := 0; i < retry; i++ {
		resp, err := http.Get(url)
		if err != nil {
			time.Sleep(2)
			continue
		}
		if resp.StatusCode == http.StatusOK {
			return
		}
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}

func get(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte("hello"))
	w.WriteHeader(http.StatusOK)
}
