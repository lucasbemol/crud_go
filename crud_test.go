package main
//@author Lucas Marins Ramos

import (
		"bytes"
		"io/ioutil"
		"net/http"
		"testing"
		"log"
		"encoding/json"
		"strconv"

)

 var produt_id = 0

//Define objeto Product(Para o resultado)
type Product struct {
	Results []struct {
		Id_product int `json:"id_product"`
		Name string `json:"name" binding:"required"`
		Price float32 `json:"price"`
		Expiration_date  string `json:"expiration_date"`
	} `json:"results"`
	Status string `json:"status"`
}

//Test add POST
func TestCreate(t *testing.T) {
	resp, err := http.Post(
		"http://localhost:3000/product/add/data.json",
		"application/json",
		bytes.NewReader([]byte(`{"name":"testProdutoLucas",
					 "price": 10,
					 "expiration_date": "2016/11/23"}`)))
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}
	if string(body) != "{\"status\":\"success\"}" {
		t.Error(string(body))
		t.Error("Update error")
	}
}

//test searchByName GET
func TestQuery(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/product/searchByName/testProdutoLucas")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	pro := Product{}
	err = json.Unmarshal([]byte(string(body)), &pro)

	if err != nil {
		t.Error(err)
	}
	 if pro.Results[0].Name != "testProdutoLucas" {
	 	t.Error("Query error")
	}

	produt_id = pro.Results[0].Id_product
}

//Test edit PUT
func TestUpdate(t *testing.T) {
	req, err := http.NewRequest(
		"PUT",
		"http://localhost:3000/product/"+strconv.Itoa(produt_id)+"/data.json",
		bytes.NewReader([]byte(`{"name":"update_name",
					 "price": 55.4,
					 "expiration_date":"2015-04-22 08:00:00"}`)))
	if err != nil {
		t.Error(err)
	}
	defer req.Body.Close()
	
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

//Test delete DELETE
func TestDelete(t *testing.T) {
	req, err := http.NewRequest("DELETE", "http://localhost:3000/product/delete/"+strconv.Itoa(produt_id), nil)
	if err != nil {
		t.Error("*********************************************")
		t.Error(err)
	}

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