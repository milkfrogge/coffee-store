package model

import "time"

type Product struct {
	Id          string
	Name        string
	Description string
	Price       int64
	Count       int64
	Category    Category
	Pics        []string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

type Category struct {
	Id   string
	Name string
}
