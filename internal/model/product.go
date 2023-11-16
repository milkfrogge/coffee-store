package model

import "time"

type Product struct {
	Id            string
	Name          string
	Description   string
	Price         uint64
	Count         uint64
	BaristaNeeded bool
	KitchenNeeded bool
	Category      Category
	Pics          []string
	CreatedAt     time.Time
}

type Category struct {
	Id   string
	Name string
}

type AddCountToProductDTO struct {
	Id    string
	Count uint64
}

type CreateProductDTO struct {
	Name          string
	Description   string
	Price         uint64
	Count         uint64
	CategoryId    string
	BaristaNeeded bool
	KitchenNeeded bool
	Pics          []string
}

type CreateCategoryDTO struct {
	Name string
}
