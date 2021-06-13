package models

// Country model for the API
type Country struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Country string `json:"country" gorm:"unique"`
}
