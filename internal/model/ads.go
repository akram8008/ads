package model

type Ads struct {
	ID          uint64   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Links       []string `json:"links"`
}
