package models

import "time"

const (
	AdsDescriptionField = "description"
	AdsLinksField       = "links"
)

type Ads struct {
	ID          uint64    `gorm:"primary_key"`
	Name        string    `gorm:"type: varchar(201);"`
	Description string    `gorm:"type: varchar(1001);"`
	Links       string    `gorm:"type: varchar(1001);"`
	Price       float64   `gorm:"type: real;"`
	CreatedDate time.Time `gorm:"type: timestamp;default:current_timestamp"`
}

/*
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Links       []string  `json:"links"`
	Price       float64   `json:"price"`
	CreatedDate time.Time `json:"created_date"`
*/
