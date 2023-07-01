package fio_connector

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const API_URL = "https://rest.fnar.net/"

func get_fio(endpoint string) (*http.Response, error) {
	fmt.Println("Requesting: " + API_URL + endpoint)
	return http.Get(API_URL + endpoint)
}

// Makes GET request to FIO endpoint and then tries to decode JSON into selected
// model
func get_fio_object(endpoint string, target interface{}) {
	resp, err := get_fio(endpoint)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Get_cx_data() []CXData {
	var data []CXData
	get_fio_object("exchange/full", &data)
	return data
}

func Get_all_recipes() []Recipe {
	var data []Recipe
	get_fio_object("recipes/allrecipes", &data)
	return data
}

func Get_cx_stations() []CXStation {
	var data []CXStation
	get_fio_object("exchange/station", &data)
	return data
}
