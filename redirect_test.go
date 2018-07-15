package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
)

func TestRedirectHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "http://", nil)
	if err != nil {
			t.Fatal(err)
	}
	req.Header.Add("X-Forwarded-Proto", "http")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RedirectHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusPermanentRedirect {
			t.Errorf("Expected status code %v, but got %v",
				http.StatusPermanentRedirect, rr.Code)
	}

	if loc := rr.Header().Get("Location"); !strings.HasPrefix(loc, "https") {
		t.Errorf("Expected https redirection, but got %v", loc)
	}
}

func TestRedirectHandlerNonHttp(t *testing.T) {
	req, err := http.NewRequest("GET", "http://", nil)
	if err != nil {
			t.Fatal(err)
	}
	req.Header.Add("X-Forwarded-Proto", "https")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RedirectHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %v, but got %v",
				http.StatusPermanentRedirect, rr.Code)
	}
}

func TestRedirectHandlerNoForwardedProto(t *testing.T) {
	req, err := http.NewRequest("GET", "http://", nil)
	if err != nil {
			t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RedirectHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %v, but got %v",
			http.StatusPermanentRedirect, rr.Code)
	}
}

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
			t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %v, but got %v",
			http.StatusOK, rr.Code)
	}
}
