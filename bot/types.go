package main

type Filter struct {
	FilterId int64  `json:"filter_id"`
	Name     string `json:"name"`
}

type Advertise struct {
	AdvertiseId int64  `json:"advertise_id"`
	Link        string `json:"link"`
	Name        string `json:"name"`
}
