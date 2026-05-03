package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Filter struct {
	FilterId int64  `json:"filter_id"`
	Name     string `json:"name"`
}

type Advertise struct {
	AdvertiseId int64  `json:"advertise_id"`
	Link        string `json:"link"`
	Name        string `json:"name"`
}

func getAllFilters() (filters []Filter) {
	resp, _ := http.Get(BASE_URL + "filters/")
	json.NewDecoder(resp.Body).Decode(&filters)
	return
}

func getFilteredAdvertises(filtersIds []int64) (advertises []Advertise) {
	query := BASE_URL + "advertises/?"
	for _, filterId := range filtersIds {
		query += "filter=" + strconv.FormatInt(filterId, 10) + "&" // filter=x&filter=y&...
	}
	resp, _ := http.Get(query[:len(query)-1]) // Without last "&"
	json.NewDecoder(resp.Body).Decode(&advertises)

	return
}
