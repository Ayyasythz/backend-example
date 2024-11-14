package wardrobe

import (
	"context"
	"github.com/google/uuid"
	"sagara_backend_test/internal/domain/model"
	"sagara_backend_test/internal/usecases/request"
	"sagara_backend_test/internal/usecases/response"
	"sagara_backend_test/lib/log"
	"sagara_backend_test/lib/tracing"
)

func (m *Module) AddStock(ctx context.Context, id *uuid.UUID, add int) (*response.WardrobeResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "WardrobeUseCases.AddStock")
	defer span.End()

	existingWardrobe, err := m.wardrobeRepo.GetById(ctx, id)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id":    id,
		}).ErrorWithCtx(ctx, "[WardrobeUseCases.AddStock] Failed to get wardrobe by ID")
		return nil, err
	}

	stockNow := existingWardrobe.Stock + add

	err = m.wardrobeRepo.AddStock(ctx, id, stockNow)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id":    id,
		}).ErrorWithCtx(ctx, "[WardrobeUseCases.AddStock] Failed to add stock")
		return nil, err
	}

	wardrobeResponse := &response.WardrobeResponse{
		ID:    existingWardrobe.ID.String(),
		Name:  existingWardrobe.Name,
		Color: existingWardrobe.Color,
		Size:  existingWardrobe.Size,
		Price: existingWardrobe.Price,
		Stock: stockNow,
	}

	return wardrobeResponse, nil
}

func (m *Module) SubStock(ctx context.Context, id *uuid.UUID, def int) (*response.WardrobeResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "WardrobeUseCases.SubStock")
	defer span.End()

	existingWardrobe, err := m.wardrobeRepo.GetById(ctx, id)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id":    id,
		}).ErrorWithCtx(ctx, "[WardrobeUseCases.SubStock] Failed to get wardrobe by ID")
		return nil, err
	}

	stockNow := existingWardrobe.Stock - def

	err = m.wardrobeRepo.SubStock(ctx, id, stockNow)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id":    id,
		}).ErrorWithCtx(ctx, "[WardrobeUseCases.SubStock] Failed to add stock")
		return nil, err
	}

	wardrobeResponse := &response.WardrobeResponse{
		ID:    existingWardrobe.ID.String(),
		Name:  existingWardrobe.Name,
		Color: existingWardrobe.Color,
		Size:  existingWardrobe.Size,
		Price: existingWardrobe.Price,
		Stock: stockNow,
	}

	return wardrobeResponse, nil
}

func (m *Module) Search(ctx context.Context, color, size string) (*[]response.WardrobeResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "WardrobeUseCases.Search")
	defer span.End()

	wardrobes, err := m.wardrobeRepo.Search(ctx, color, size)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[WardrobeUseCases.Search] Failed to get wardrobes")
		return nil, err
	}

	var wardrobeResponses []response.WardrobeResponse
	for _, wardrobe := range *wardrobes {
		wardrobeResponses = append(wardrobeResponses, response.WardrobeResponse{
			ID:    wardrobe.ID.String(),
			Name:  wardrobe.Name,
			Color: wardrobe.Color,
			Size:  wardrobe.Size,
			Price: wardrobe.Price,
			Stock: wardrobe.Stock,
		})
	}

	return &wardrobeResponses, nil

}

func (m *Module) GetAllWardrobe(ctx context.Context) (*[]response.WardrobeResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "WardrobeUseCases.GetAllWardrobe")
	defer span.End()

	wardrobes, err := m.wardrobeRepo.GetAll(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[WardrobeUseCases.GetAllWardrobe] Failed to get all wardrobe")
		return nil, err
	}

	var wardrobeResponses []response.WardrobeResponse
	for _, wardrobe := range *wardrobes {
		wardrobeResponses = append(wardrobeResponses, response.WardrobeResponse{
			ID:    wardrobe.ID.String(),
			Name:  wardrobe.Name,
			Color: wardrobe.Color,
			Size:  wardrobe.Size,
			Price: wardrobe.Price,
			Stock: wardrobe.Stock,
		})
	}

	return &wardrobeResponses, nil
}

func (m *Module) GetWardrobe(ctx context.Context, id *uuid.UUID) (*response.WardrobeResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "WardrobeUseCases.GetWardrobe")
	defer span.End()

	wardrobe, err := m.wardrobeRepo.GetById(ctx, id)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id":    id,
		}).ErrorWithCtx(ctx, "[WardrobeUseCases.GetAllWardrobe] Failed to get all wardrobe")
		return nil, err
	}

	return &response.WardrobeResponse{
		ID:    wardrobe.ID.String(),
		Name:  wardrobe.Name,
		Color: wardrobe.Color,
		Size:  wardrobe.Size,
		Price: wardrobe.Price,
		Stock: wardrobe.Stock,
	}, nil
}

func (m *Module) InsertWardrobe(ctx context.Context, request *request.WardrobeInsertRequest) (*response.WardrobeResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "WardrobeUseCases.InsertWardrobe")
	defer span.End()

	newWardrobe := &model.Wardrobe{
		ID:    uuid.New(),
		Name:  request.Name,
		Color: request.Color,
		Size:  request.Size,
		Price: request.Price,
		Stock: request.Stock,
	}

	err := m.wardrobeRepo.Insert(ctx, newWardrobe)
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err,
			"request": request,
		}).ErrorWithCtx(ctx, "[WardrobeUseCases.InsertWardrobe] Failed to insert wardrobe")
		return nil, err
	}

	wardrobeResponse := &response.WardrobeResponse{
		ID:    newWardrobe.ID.String(),
		Name:  newWardrobe.Name,
		Color: newWardrobe.Color,
		Size:  newWardrobe.Size,
		Price: newWardrobe.Price,
		Stock: newWardrobe.Stock,
	}

	return wardrobeResponse, nil
}

func (m *Module) UpdateWardrobe(ctx context.Context, id *uuid.UUID, request *request.WardrobeUpdateRequest) (*response.WardrobeResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "WardrobeUseCases.UpdateWardrobe")
	defer span.End()

	existingWardrobe, err := m.wardrobeRepo.GetById(ctx, id)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id":    id,
		}).ErrorWithCtx(ctx, "[WardrobeUseCases.UpdateWardrobe] Failed to get wardrobe by ID")
		return nil, err
	}

	existingWardrobe.Name = request.Name
	existingWardrobe.Color = request.Color
	existingWardrobe.Size = request.Size
	existingWardrobe.Price = request.Price
	existingWardrobe.Stock = request.Stock

	err = m.wardrobeRepo.Update(ctx, existingWardrobe)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id":    id,
		}).ErrorWithCtx(ctx, "[WardrobeUseCases.UpdateWardrobe] Failed to update wardrobe")
		return nil, err
	}

	wardrobeResponse := &response.WardrobeResponse{
		ID:    existingWardrobe.ID.String(),
		Name:  existingWardrobe.Name,
		Color: existingWardrobe.Color,
		Size:  existingWardrobe.Size,
		Price: existingWardrobe.Price,
		Stock: existingWardrobe.Stock,
	}

	return wardrobeResponse, nil
}

func (m *Module) DeleteWardrobe(ctx context.Context, id *uuid.UUID) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "WardrobeUseCases.DeleteWardrobe")
	defer span.End()

	err := m.wardrobeRepo.Delete(ctx, id)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id":    id,
		}).ErrorWithCtx(ctx, "[WardrobeUseCases.DeleteWardrobe] Failed to delete wardrobe")
		return err
	}

	return nil
}

func (m *Module) GetAvailable(ctx context.Context) (*[]response.WardrobeResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "WardrobeUseCases.GetAvailable")
	defer span.End()

	wardrobes, err := m.wardrobeRepo.GetAvailable(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[WardrobeUseCases.GetAvailable] Failed to get available wardrobe")
		return nil, err
	}

	var wardrobeResponses []response.WardrobeResponse
	for _, wardrobe := range *wardrobes {
		wardrobeResponses = append(wardrobeResponses, response.WardrobeResponse{
			ID:    wardrobe.ID.String(),
			Name:  wardrobe.Name,
			Color: wardrobe.Color,
			Size:  wardrobe.Size,
			Price: wardrobe.Price,
			Stock: wardrobe.Stock,
		})
	}

	return &wardrobeResponses, nil
}

func (m *Module) GetUnavailable(ctx context.Context) (*[]response.WardrobeResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "WardrobeUseCases.GetUnavailable")
	defer span.End()

	wardrobes, err := m.wardrobeRepo.GetUnavailable(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[WardrobeUseCases.GetAvailable] Failed to get unavailable wardrobe")
		return nil, err
	}

	var wardrobeResponses []response.WardrobeResponse
	for _, wardrobe := range *wardrobes {
		wardrobeResponses = append(wardrobeResponses, response.WardrobeResponse{
			ID:    wardrobe.ID.String(),
			Name:  wardrobe.Name,
			Color: wardrobe.Color,
			Size:  wardrobe.Size,
			Price: wardrobe.Price,
			Stock: wardrobe.Stock,
		})
	}

	return &wardrobeResponses, nil
}

func (m *Module) GetLessThan(ctx context.Context, amount int) (*[]response.WardrobeResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "WardrobeUseCases.GetLessThan")
	defer span.End()

	wardrobes, err := m.wardrobeRepo.GetLessThan(ctx, amount)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[WardrobeUseCases.GetLessThan] Failed to get wardrobe")
		return nil, err
	}

	var wardrobeResponses []response.WardrobeResponse
	for _, wardrobe := range *wardrobes {
		wardrobeResponses = append(wardrobeResponses, response.WardrobeResponse{
			ID:    wardrobe.ID.String(),
			Name:  wardrobe.Name,
			Color: wardrobe.Color,
			Size:  wardrobe.Size,
			Price: wardrobe.Price,
			Stock: wardrobe.Stock,
		})
	}

	return &wardrobeResponses, nil
}
