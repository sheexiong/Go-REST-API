package models

// Property model for the API
type Property struct {
	ID           uint    `json:"id" gorm:"primary_key"`
	PostType     string  `json:"post_type"` // rent or sale
	Price        float64 `json:"price"`
	Desc         string  `json:"desc"`
	PropertyType string  `json:"property_type"` //hdb, condo or landard
	UnitSize     int     `json:"unit_size"`
	Addr         string  `json:"addr"`
	Bedrooms     int     `json:"bedrooms"`
	Bathrooms    int     `json:"bathrooms"`
	Country      string  `json:"country"`
}
