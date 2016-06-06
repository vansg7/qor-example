package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/qor/publish"
	"github.com/qor/slug"
	"github.com/qor/sorting"
)

type Product struct {
	gorm.Model
	Name         string
	NameWithSlug slug.Slug `l10n:"sync"`
	Code         string    `l10n:"sync"`
	Prices       []Price

	publish.Status
	sorting.SortingDESC
}

type ProductVariation struct {
	gorm.Model
	SKU       string
	Prices    []Price
	ProductID *Product
	Product   *Product
}

type Price struct {
	gorm.Model
	Amount         uint
	OriginalAmount uint
	StartDate      *time.Time
	EndDate        *time.Time
}
