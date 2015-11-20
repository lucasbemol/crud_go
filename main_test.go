package main


import (
		"bytes"
		"io/ioutil"
		"net/http"
		"testing"
)

func TestCreate(t *testing.T) {
	resp, err := http.Post(
		"http://localhost:3000/product/add/data.json",
		"application/json",
		bytes.NewReader([]byte(`{"name":"new_des",
					 "price":"10.00",
					 "expiration_date": "2018-01-01 08:00:00"}`)))
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	if string(body) != "{\"status\":\"success\"}" {
		t.Error(string(body))
		t.Error("Update error")
	}
}