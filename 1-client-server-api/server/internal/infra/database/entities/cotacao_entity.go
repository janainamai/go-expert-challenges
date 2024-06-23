package entities

import "gorm.io/gorm"

type (
	CotacaoEntity struct {
		ID         int `gorm:"primaryKey"`
		Code       string
		Codein     string
		Name       string
		High       string
		Low        string
		VarBid     string
		PctChange  string
		Bid        string
		Ask        string
		Timestamp  string
		CreateDate string
		gorm.Model
	}
)
