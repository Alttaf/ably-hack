package main

import (
	"bytes"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestHello(t *testing.T) {
	os.Clearenv()

	r := httptest.NewRequest("GET", "/hello/sd", bytes.NewReader([]byte{}))
	w := httptest.NewRecorder()

	var p httprouter.Params = []httprouter.Param{
		{Key: "name", Value: "name_value"},
	}

	Hello(w, r, p)

	wanted := `ABLY_API_KEY unset. This must be set from an .env file`
	if strings.TrimSpace(w.Body.String()) != wanted {
		t.Errorf(`wanted body "%v" got "%v"`, wanted, w.Body.String())
	}
}

func TestIndex(t *testing.T) {
	r := httptest.NewRequest("GET", "/", bytes.NewReader([]byte{}))
	w := httptest.NewRecorder()

	var p []httprouter.Param

	Index(w, r, p)

	if w.Code != 200 {
		t.Errorf("wanted response code 200, got %v, body: %s", w.Code, w.Body)
	}

	wanted := `Welcome!`
	if strings.TrimSpace(w.Body.String()) != wanted {
		t.Errorf(`wanted body "%v" got "%v"`, wanted, w.Body.String())
	}
}
