package request

import (
	"sagara_backend_test/lib/custerr"
	"sagara_backend_test/lib/response"
	"sagara_backend_test/pkg/constants"
	"sagara_backend_test/pkg/constants/errorcode"
)

type WardrobeInsertRequest struct {
	Name  string  `json:"name"`
	Color string  `json:"color"`
	Size  string  `json:"size"`
	Price float32 `json:"price"`
	Stock int     `json:"stock"`
}

type WardrobeUpdateRequest struct {
	Name  string  `json:"name"`
	Color string  `json:"color"`
	Size  string  `json:"size"`
	Price float32 `json:"price"`
	Stock int     `json:"stock"`
}

type WardrobeAddSubRequest struct {
	Amount int `json:"amount"`
}

func (w *WardrobeUpdateRequest) ValidateUpdateWardrobe() error {
	if w.Name == constants.EmptyString {
		return custerr.ErrChain{
			Message: errorcode.NameEmpty.Message,
			Code:    errorcode.NameEmpty.Code,
			Type:    response.ErrBadRequest,
		}
	}
	if w.Color == constants.EmptyString {
		return custerr.ErrChain{
			Message: errorcode.ColorEmpty.Message,
			Code:    errorcode.ColorEmpty.Code,
			Type:    response.ErrBadRequest,
		}
	}
	if w.Size == constants.EmptyString {
		return custerr.ErrChain{
			Message: errorcode.SizeEmpty.Message,
			Code:    errorcode.SizeEmpty.Code,
			Type:    response.ErrBadRequest,
		}
	}

	return nil
}
