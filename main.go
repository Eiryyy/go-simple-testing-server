package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

type s struct {
	Counter int `json:"counter"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	u, _ := url.Parse(r.RequestURI)
	counter, _ := strconv.Atoi(u.Query().Get("c"))
	j := s{Counter: counter}
	json, _ := json.Marshal(j)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
