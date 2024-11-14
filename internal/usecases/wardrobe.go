package usecases

import (
	"context"
	"github.com/google/uuid"
	"sagara_backend_test/internal/usecases/request"
	"sagara_backend_test/internal/usecases/response"
)

type WardrobeUseCases interface {
	GetAllWardrobe(ctx context.Context) (*[]response.WardrobeResponse, error)
	GetWardrobe(ctx context.Context, id *uuid.UUID) (*response.WardrobeResponse, error)
	InsertWardrobe(ctx context.Context, request *request.WardrobeInsertRequest) (*response.WardrobeResponse, error)
	UpdateWardrobe(ctx context.Context, id *uuid.UUID, request *request.WardrobeUpdateRequest) (*response.WardrobeResponse, error)
	DeleteWardrobe(ctx context.Context, id *uuid.UUID) error
	Search(ctx context.Context, color, size string) (*[]response.WardrobeResponse, error)
	AddStock(ctx context.Context, id *uuid.UUID, add int) (*response.WardrobeResponse, error)
	SubStock(ctx context.Context, id *uuid.UUID, def int) (*response.WardrobeResponse, error)
	GetAvailable(ctx context.Context) (*[]response.WardrobeResponse, error)
	GetUnavailable(ctx context.Context) (*[]response.WardrobeResponse, error)
	GetLessThan(ctx context.Context, amount int) (*[]response.WardrobeResponse, error)
}
