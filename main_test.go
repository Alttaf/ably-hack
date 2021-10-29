package main

import (
	"bytes"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestHello(t *testing.T) {
	r := httptest.NewRequest("GET", "/hello/sd", bytes.NewReader([]byte{}))
	w := httptest.NewRecorder()

	var p httprouter.Params = []httprouter.Param{
		{Key: "name", Value: "name_value"},
	}

	Hello(w, r, p)

	if w.Code != 200 {
		t.Errorf("wanted response code 200, got %v, body: %s", w.Code, w.Body)
	}

	wanted := `Hello Ably, name_value!`
	if strings.TrimSpace(w.Body.String()) != wanted {
		t.Errorf(`wanted body "%v" got "%v"`, wanted, w.Body.String())
	}
}

func TestIndex(t *testing.T) {

}
