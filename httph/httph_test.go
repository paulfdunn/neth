package httph

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPBodyUnmarshal(t *testing.T) {
	id := "this is a test struct"
	type testStruct struct {
		ID string
	}
	tsIn := testStruct{ID: id}
	tsOut := testStruct{}
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		HTTPBodyUnmarshal(w, r, &tsOut)
	}))
	defer testServer.Close()

	b, err := json.Marshal(tsIn)
	if err != nil {
		t.Errorf("marshal error: %v", err)
		return
	}
	resp, err := http.Post(testServer.URL, "application/json", bytes.NewBuffer(b))
	if err != nil {
		t.Errorf("POST error: %v", err)
		return
	}
	if resp.StatusCode != 200 {
		t.Errorf("status code: %d", resp.StatusCode)
		return
	}
	if tsOut.ID != id {
		t.Errorf("incorrect tsOut: %+v", tsOut)
		return
	}

	resp, err = http.Post(testServer.URL, "application/json", bytes.NewBuffer(b[1:]))
	if err != nil {
		t.Errorf("POST error: %v", err)
		return
	}
	if resp.StatusCode != http.StatusUnprocessableEntity {
		t.Errorf("status code: %d", resp.StatusCode)
		return
	}
}
