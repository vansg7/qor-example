package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ProductVariation struct {
	gorm.Model
}

type Price struct {
	gorm.Model
	Amount         uint
	OriginalAmount uint
	StartDate      *time.Time
	EndDate        *time.Time
}
