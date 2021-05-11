package main

type Language struct {
	Name       string   `json:"name"`
	Mode       string   `json:"mode"`
	Mimes      []string `json:"mimes"`
	Extensions []string `json:"ext"`
	Color      string   `json:"color"`
}