package model

import (
	"gorm.io/gorm"
)

type Restaurant struct {
	gorm.Model
	Name          string   `json:"name"`
	Address       string   `json:"address"`
	HasPickup     bool     `json:"has_pickup"`
	HasDelivery   bool     `json:"has_delivery"`
	IsVegeterian  bool     `json:"is_vegeterian"`
	IsVegan       bool     `json:"is_vegan"`
	IsPescaterian bool     `json:"is_pescaterian"`
	CuisineTags   []string `json:"cuisines_tags"`
	AverageRating float32  `json:"average_rating"`
	Reviews       string   `json:"reviews"`
	Hours         string   `json:"hours"`
	HasOrdered    bool     `json:"has_ordered"`
}
