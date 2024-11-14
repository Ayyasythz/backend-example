package repository

import (
	"context"
	"github.com/google/uuid"
	"sagara_backend_test/internal/domain/model"
)

type WardrobeRepository interface {
	Insert(ctx context.Context, wardrobe *model.Wardrobe) error
	Update(ctx context.Context, wardrobe *model.Wardrobe) error
	GetAll(ctx context.Context) (*[]model.Wardrobe, error)
	GetById(ctx context.Context, id *uuid.UUID) (*model.Wardrobe, error)
	Delete(ctx context.Context, id *uuid.UUID) error
	Search(ctx context.Context, color, size string) (*[]model.Wardrobe, error)
	AddStock(ctx context.Context, id *uuid.UUID, addition int) error
	SubStock(ctx context.Context, id *uuid.UUID, def int) error
	GetAvailable(ctx context.Context) (*[]model.Wardrobe, error)
	GetUnavailable(ctx context.Context) (*[]model.Wardrobe, error)
	GetLessThan(ctx context.Context, amount int) (*[]model.Wardrobe, error)
}
