package models

import "time"

type ComboCategory struct{
	ComboCategoryName	string `json:"combo_category_name"`
	ComboCategoryID		int `json:"combo_category_id"`
	Created time.Time
	Modified time.Time
}


type Combo struct {
	ComboID int `json:"combo_id"`
	ComboCost float64 `json:"combo_cost"`
	ComboValidity string `json:"combo_validity"`
	ComboCategoryID int `json:"combo_category_id"`
	ComboBundles []Bundles `json:"combo_bundles"`
	Created time.Time
	Modified time.Time
}

// ComboCategories is an array of ComboCategories
type ComboCategories []ComboCategory

// Combos is an array of Combos
type Combos []Combo
