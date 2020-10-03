package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHpaExample(t *testing.T) {

	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	want := "OK!"
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)
	handler.ServeHTTP(w, r)

	resp := w.Result()

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("Unexpected error %d", err)
		}
		bodyString := string(bodyBytes)
		if got := bodyString; got != want {
			t.Errorf("handler() = %q, want %q", got, want)
		}
	} else if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code %d", resp.StatusCode)
	}
}

func TestLiveness(t *testing.T) {

	r, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}

	want := "ok"
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(status)
	handler.ServeHTTP(w, r)

	resp := w.Result()

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("Unexpected error %d", err)
		}
		bodyString := string(bodyBytes)
		if got := bodyString; got != want {
			t.Errorf("handler() = %q, want %q", got, want)
		}
	} else if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code %d", resp.StatusCode)
	}
}

func TestReadiness(t *testing.T) {

	r, err := http.NewRequest("GET", "/readyz", nil)
	if err != nil {
		t.Fatal(err)
	}

	want := "ok"
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(status)
	handler.ServeHTTP(w, r)

	resp := w.Result()

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("Unexpected error %d", err)
		}
		bodyString := string(bodyBytes)
		if got := bodyString; got != want {
			t.Errorf("handler() = %q, want %q", got, want)
		}
	} else if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code %d", resp.StatusCode)
	}
}
