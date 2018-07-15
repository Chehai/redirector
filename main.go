package main

import (
	"net/http"
	"flag"
	"log"
	"strconv"
	"net/url"
)

func RedirectHandler(w http.ResponseWriter, req *http.Request) {
	if req.Header.Get("X-Forwarded-Proto") != "http" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	target := url.URL{
		Scheme: "https",
		Host: req.Host,
		Path: req.URL.Path,
		RawQuery: req.URL.RawQuery,
	}
	http.Redirect(w, req, target.String(), http.StatusPermanentRedirect)
}

func HealthCheckHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	port := flag.Int("port", 8888, "port")
	flag.Parse()

	http.HandleFunc("/", RedirectHandler)
	http.HandleFunc("/healthz", HealthCheckHandler)

	lisentAddr := ":" + strconv.Itoa(*port)
	log.Fatal(http.ListenAndServe(lisentAddr, nil))
}
