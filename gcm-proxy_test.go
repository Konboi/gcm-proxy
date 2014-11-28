package gcm

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestReciver(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	Reciver(w, req)
	if w.Code == http.StatusOK {
		t.Fatalf("Allowd only POST Method but expected status 200; received %d", w.Code)
	}

	data := url.Values{}
	data.Add("token", "")

	req, err = http.NewRequest("POST", "/", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w = httptest.NewRecorder()
	Reciver(w, req)
	if w.Code == http.StatusOK {
		t.Fatalf("token is empty but expected status 200; received %d", w.Code)
	}

	data = url.Values{}
	data.Add("token", "abcdef,ghijk,12345")
	data.Add("alert", "hoge")

	req, err = http.NewRequest("POST", "/", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w = httptest.NewRecorder()
	Reciver(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("Not return 200; received %d", w.Code)
	}

	data = url.Values{}
	data.Add("token", "abcdef,ghijk,12345")
	data.Add("payload", `{"message": "test"}`)
	req, err = http.NewRequest("POST", "/", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w = httptest.NewRecorder()
	Reciver(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("Not return 200; received %d", w.Code)
	}
}

func Testsend(t *testing.T) {}
