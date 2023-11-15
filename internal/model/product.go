package model

import "time"

type Product struct {
	Id          string
	Name        string
	Description string
	Price       uint64
	Count       uint64
	Category    Category
	Pics        []string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
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
	Name        string
	Description string
	Price       uint64
	Count       uint64
	CategoryId  string
	Pics        []string
}

type CreateCategoryDTO struct {
	Name string
}
