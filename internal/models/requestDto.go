package models

type RequestDto struct {
	URL  string `json:"url"`
	Hash string `json:"hash"`
	Id   string `json:"id"`
}
