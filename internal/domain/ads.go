package domain

type AdsRequest struct {
	Name        *string   `json:"name"`
	Description *string   `json:"description"`
	Links       *[]string `json:"links"`
	Price       *float64  `json:"price"`
}

type AdsResponse struct {
	Name        *string   `json:"name,omitempty"`
	MainLink    *string   `json:"link,omitempty"`
	Price       *float64  `json:"price,omitempty"`
	Description *string   `json:"description,omitempty"`
	Links       *[]string `json:"links,omitempty"`
}
