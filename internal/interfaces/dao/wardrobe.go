package dao

import (
	"context"
	sql2 "database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"sagara_backend_test/internal/domain/model"
	"sagara_backend_test/internal/domain/repository"
	"sagara_backend_test/lib/database/sql"
	"sagara_backend_test/lib/log"
	"sagara_backend_test/lib/tracing"
	"sagara_backend_test/lib/txmanager/utils"
	"strings"
	"time"
)

type WardrobeRepository struct {
	db *sql.Store
}

type OptsWardrobeRepository struct {
	DB *sql.Store
}

const (
	insertWardrobe    = `INSERT INTO wardrobe (id, name, color, size, price, stock, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	selectWardrobe    = `SELECT id, name, color, size, price, stock, created_at, updated_at FROM wardrobe WHERE TRUE %s`
	selectAllWardrobe = `SELECT id, name, color, size, price, stock, created_at, updated_at FROM wardrobe`
	updateWardrobe    = `UPDATE wardrobe SET %s WHERE TRUE %s`
	deleteWardrobe    = `DELETE FROM wardrobe WHERE TRUE %s`
)

func NewWardrobeRepository(opts *OptsWardrobeRepository) repository.WardrobeRepository {
	return &WardrobeRepository{db: opts.DB}
}

func (w *WardrobeRepository) Insert(ctx context.Context, wardrobe *model.Wardrobe) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "WardrobeRepository.Insert")
	defer span.End()

	var (
		err error
	)

	sqlTrx := utils.GetSqlTx(ctx)
	if sqlTrx != nil {
		_, err = sqlTrx.ExecContext(ctx, insertWardrobe, wardrobe.ID, wardrobe.Name, wardrobe.Color,
			wardrobe.Size, wardrobe.Price, wardrobe.Stock, wardrobe.CreatedAt, wardrobe.UpdatedAt)
	} else {
		_, err = w.db.GetMaster().ExecContext(ctx, insertWardrobe, wardrobe.ID, wardrobe.Name, wardrobe.Color,
			wardrobe.Size, wardrobe.Price, wardrobe.Stock, wardrobe.CreatedAt, wardrobe.UpdatedAt)
	}

	if err != nil {
		if pqErr, valid := err.(*pq.Error); valid {
			switch pqErr.Code {
			case "23505":
				log.WithFields(log.Fields{
					"error":    err,
					"wardrobe": *wardrobe,
				}).ErrorWithCtx(ctx, "[WardrobeRepository.Insert] Duplicate Entry")
				return ErrDuplicate
			}
		}
		log.WithFields(log.Fields{
			"error":    err,
			"wardrobe": *wardrobe,
		}).ErrorWithCtx(ctx, "[WardrobeRepository.Insert] Failed to Insert")
		return err
	}

	return nil
}

func (w *WardrobeRepository) Update(ctx context.Context, wardrobe *model.Wardrobe) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "WardrobeRepository.Update")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)

	var (
		args []any
		err  error
	)

	setQuery := "name = $1, color = $2, size = $3, price = $4, stock = $5, updated_at = $6"
	whereQuery := " AND id = $7"
	args = append(args, wardrobe.Name, wardrobe.Color, wardrobe.Size, wardrobe.Price, wardrobe.Stock, time.Now(), wardrobe.ID)

	query := fmt.Sprintf(updateWardrobe, setQuery, whereQuery)

	if sqlTrx != nil {
		_, err = sqlTrx.ExecContext(ctx, query, args...)
	} else {
		_, err = w.db.GetMaster().ExecContext(ctx, query, args...)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error":    err,
			"wardrobe": wardrobe,
		}).ErrorWithCtx(ctx, "[WardrobeRepository.Update] Failed to update wardrobe")
		return err
	}
	return nil
}

func (w *WardrobeRepository) GetAll(ctx context.Context) (*[]model.Wardrobe, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "WardrobeRepository.GetAll")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)

	var (
		wardrobe []model.Wardrobe
		err      error
	)

	if sqlTrx != nil {
		err = sqlTrx.SelectContext(ctx, &wardrobe, selectAllWardrobe)
	} else {
		err = w.db.GetMaster().SelectContext(ctx, &wardrobe, selectAllWardrobe)
	}

	if err != nil {
		if errors.Is(err, sql2.ErrNoRows) {
			return nil, ErrNoResult
		}
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[WardrobeRepository.GetAll] Failed to get all wardrobe")
		return nil, err
	}

	return &wardrobe, nil
}

func (w *WardrobeRepository) GetById(ctx context.Context, id *uuid.UUID) (*model.Wardrobe, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "WardrobeRepository.GetByID")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)

	var (
		args     []any
		wardrobe model.Wardrobe
		err      error
	)

	whereQuery := " AND id = $1"
	args = append(args, id)

	query := fmt.Sprintf(selectWardrobe, whereQuery)

	if sqlTrx != nil {
		err = sqlTrx.GetContext(ctx, &wardrobe, query, args...)
	} else {
		err = w.db.GetMaster().GetContext(ctx, &wardrobe, query, args...)
	}

	if err != nil {
		if errors.Is(err, sql2.ErrNoRows) {
			return nil, ErrNoResult
		}
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[WardrobeRepository.GetById] Failed to get wardrobe by id")
		return nil, err
	}

	return &wardrobe, nil
}

func (w *WardrobeRepository) Delete(ctx context.Context, id *uuid.UUID) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "WardrobeRepository.Delete")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)

	var (
		args []any
		err  error
	)

	whereQuery := " AND id = $1"
	args = append(args, id)

	query := fmt.Sprintf(deleteWardrobe, whereQuery)

	if sqlTrx != nil {
		_, err = sqlTrx.ExecContext(ctx, query, args...)
	} else {
		_, err = w.db.GetMaster().ExecContext(ctx, query, args...)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id":    id,
		}).ErrorWithCtx(ctx, "[WardrobeRepository.Delete] Failed to delete wardrobe")
		return err
	}
	return nil
}

func (w *WardrobeRepository) Search(ctx context.Context, color, size string) (*[]model.Wardrobe, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "WardrobeRepository.Search")
	defer span.End()

	var (
		wardrobes  []model.Wardrobe
		args       []any
		whereQuery string
		err        error
	)

	sqlTrx := utils.GetSqlTx(ctx)

	color = strings.ToLower(color)
	size = strings.ToLower(size)

	if color != "" && size != "" {
		whereQuery = " AND color = $1 AND size = $2"
		args = append(args, color, size)
	} else if color != "" {
		whereQuery = " AND color = $1"
		args = append(args, color)
	} else if size != "" {
		whereQuery = " AND size = $1"
		args = append(args, size)
	} else {
		whereQuery = ""
	}

	query := fmt.Sprintf(selectWardrobe, whereQuery)

	if sqlTrx != nil {
		err = sqlTrx.SelectContext(ctx, &wardrobes, query, args...)
	} else {
		err = w.db.GetMaster().SelectContext(ctx, &wardrobes, query, args...)
	}

	fmt.Println(query)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"color": color,
			"size":  size,
		}).ErrorWithCtx(ctx, "[WardrobeRepository.Delete] Failed to delete wardrobe")
		return nil, err
	}
	return &wardrobes, nil

}

func (w *WardrobeRepository) AddStock(ctx context.Context, id *uuid.UUID, addition int) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "WardrobeRepository.AddStock")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)

	var (
		args []any
		err  error
	)

	setQuery := "stock = $1"
	whereQuery := " AND id = $2"
	args = append(args, addition, id)

	query := fmt.Sprintf(updateWardrobe, setQuery, whereQuery)

	if sqlTrx != nil {
		_, err = sqlTrx.ExecContext(ctx, query, args...)
	} else {
		_, err = w.db.GetMaster().ExecContext(ctx, query, args...)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id":    id,
		}).ErrorWithCtx(ctx, "[WardrobeRepository.AddStock] Failed to add stock")
		return err
	}
	return nil
}

func (w *WardrobeRepository) SubStock(ctx context.Context, id *uuid.UUID, def int) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "WardrobeRepository.SubStock")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)

	var (
		args []any
		err  error
	)

	setQuery := "stock = $1"
	whereQuery := " AND id = $2"
	args = append(args, def, id)

	query := fmt.Sprintf(updateWardrobe, setQuery, whereQuery)

	if sqlTrx != nil {
		_, err = sqlTrx.ExecContext(ctx, query, args...)
	} else {
		_, err = w.db.GetMaster().ExecContext(ctx, query, args...)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id":    id,
		}).ErrorWithCtx(ctx, "[WardrobeRepository.SubStock] Failed to sub stock")
		return err
	}
	return nil
}

func (w *WardrobeRepository) GetAvailable(ctx context.Context) (*[]model.Wardrobe, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "WardrobeRepository.GetAvailable")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)

	var (
		wardrobe []model.Wardrobe
		err      error
	)

	whereQuery := " AND stock != 0"
	query := fmt.Sprintf(selectWardrobe, whereQuery)

	if sqlTrx != nil {
		err = sqlTrx.SelectContext(ctx, &wardrobe, query)
	} else {
		err = w.db.GetMaster().SelectContext(ctx, &wardrobe, query)
	}

	if err != nil {
		if errors.Is(err, sql2.ErrNoRows) {
			return nil, ErrNoResult
		}
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[WardrobeRepository.GetAvailable] Failed to get available wardrobe")
		return nil, err
	}

	return &wardrobe, nil
}

func (w *WardrobeRepository) GetUnavailable(ctx context.Context) (*[]model.Wardrobe, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "WardrobeRepository.GetUnavailable")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)

	var (
		wardrobe []model.Wardrobe
		err      error
	)

	whereQuery := " AND stock = 0"
	query := fmt.Sprintf(selectWardrobe, whereQuery)

	if sqlTrx != nil {
		err = sqlTrx.SelectContext(ctx, &wardrobe, query)
	} else {
		err = w.db.GetMaster().SelectContext(ctx, &wardrobe, query)
	}

	if err != nil {
		if errors.Is(err, sql2.ErrNoRows) {
			return nil, ErrNoResult
		}
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[WardrobeRepository.GetUnavailable] Failed to unavailable wardrobe")
		return nil, err
	}

	return &wardrobe, nil
}

func (w *WardrobeRepository) GetLessThan(ctx context.Context, amount int) (*[]model.Wardrobe, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "WardrobeRepository.GetLessThan")
	defer span.End()

	sqlTrx := utils.GetSqlTx(ctx)

	var (
		wardrobe   []model.Wardrobe
		err        error
		whereQuery string
		args       []any
	)

	if amount != 0 {
		whereQuery = " AND stock < $1 "
		args = append(args, amount)
	} else {
		whereQuery = " AND stock < 5"
	}
	query := fmt.Sprintf(selectWardrobe, whereQuery)

	if sqlTrx != nil {
		err = sqlTrx.SelectContext(ctx, &wardrobe, query, args...)
	} else {
		err = w.db.GetMaster().SelectContext(ctx, &wardrobe, query, args...)
	}

	if err != nil {
		if errors.Is(err, sql2.ErrNoRows) {
			return nil, ErrNoResult
		}
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[WardrobeRepository.GetLessThan] Failed to get wardrobe")
		return nil, err
	}

	return &wardrobe, nil
}
