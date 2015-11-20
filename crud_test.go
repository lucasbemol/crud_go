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
		bytes.NewReader([]byte(`{"name":"teste Produto",
					 "price": 10.10,
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

//test query for first record, please make sure not to delete it
func TestQuery(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/product/1")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	//verify the resp content
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	 if string(body) != "{\"results\":{\"description\":\"text\",\"due\":\"2008-01-01 10:00:00\",\"completed\":false},\"status\":\"success\"}" {
	// 	t.Error("Query error")
	}
}

func TestUpdate(t *testing.T) {
	req, err := http.NewRequest(
		"PUT",
		"http://localhost:3000/product/1/data.json",
		bytes.NewReader([]byte(`{"name":"update_name",
					 "price": 55.4,
					 "expiration_date":"2015-04-22 08:00:00"}`)))
	if err != nil {
		t.Error(err)
	}
	defer req.Body.Close()
	//verify the resp content
	resp, err := http.DefaultClient.Do(req)
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

//test for deleting 5th record in sqlite3, make sure to have at least 5 records
func TestDelete(t *testing.T) {
	//resp, err := http.Get("http://localhost:3000/task/delete/5")
	req, err := http.NewRequest("DELETE", "http://localhost:3000/product/delete/5", nil)
	if err != nil {
		t.Error("*********************************************")
		t.Error(err)
	}
	//defer req.Body.Close()

	//verify the resp content
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error("---------------------------------------------")
		t.Error(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if string(body) != "{\"status\":\"success\"}" {
		t.Error(string(body))
		t.Error("Deletion error")
	}
}