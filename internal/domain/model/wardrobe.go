package model

import "github.com/google/uuid"

type Wardrobe struct {
	BaseModel
	ID    uuid.UUID `db:"id"`
	Name  string    `db:"name"`
	Color string    `db:"color"`
	Size  string    `db:"size"`
	Price float32   `db:"price"`
	Stock int       `db:"stock"`
}
