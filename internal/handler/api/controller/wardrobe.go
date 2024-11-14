package controller

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"sagara_backend_test/internal/infrastructures/custresp"
	"sagara_backend_test/internal/usecases/request"
	"sagara_backend_test/lib/response/rest"
	"sagara_backend_test/lib/router"
	"sagara_backend_test/lib/tracing"
	"strconv"
)

// GetAll godoc
// @Summary 	Get All Wardrobe
// @Description	Get All Wardrobe
// @Tags		wardrobes
// @Accept		json
// @Produce		json
// @Success		200	{object}	jsonResponse{data=response.WardrobeResponse}
// @Router		/v1/wardrobe	[get]
func (api *API) GetAll(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetAll")
	defer span.End()

	res, err := api.wardrobeUc.GetAllWardrobe(ctx)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// Insert godoc
// @Summary 	Insert Wardrobe
// @Description	Insert Wardrobe
// @Tags		wardrobes
// @Accept		json
// @Param		wardrobes 		body 	request.WardrobeInsertRequest true "Insert Payload"
// @Produce		json
// @Success		200	{object}	jsonResponse{data=response.WardrobeResponse}
// @Router		/v1/wardrobe	[post]
func (api *API) Insert(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.Insert")
	defer span.End()

	var insertReq request.WardrobeInsertRequest
	err := json.Unmarshal(req.RawBody(), &insertReq)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	res, err := api.wardrobeUc.InsertWardrobe(ctx, &insertReq)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// Update godoc
// @Summary 	Update Wardrobe
// @Description	Update Wardrobe
// @Tags		wardrobes
// @Accept		json
// @Param		wardrobes 		body 	request.WardrobeUpdateRequest true "Update Payload"
// @Produce		json
// @Param 		id		path 		string 	false 	"wardrobe id"
// @Success		200	{object}	jsonResponse{data=response.WardrobeResponse}
// @Router		/v1/wardrobe/{id}	[put]
func (api *API) Update(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.Update")
	defer span.End()

	wardrobeIDStr := req.Params("id")
	if wardrobeIDStr == "" {
		return custresp.CustomErrorResponse(errors.New("missing id"))
	}

	wardrobeID, err := uuid.Parse(wardrobeIDStr)
	if err != nil {
		return custresp.CustomErrorResponse(errors.New("invalid id"))
	}

	var updateReq request.WardrobeUpdateRequest
	err = json.Unmarshal(req.RawBody(), &updateReq)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	err = updateReq.ValidateUpdateWardrobe()
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}
	res, err := api.wardrobeUc.UpdateWardrobe(ctx, &wardrobeID, &updateReq)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetById godoc
// @Summary 	Get Wardrobe By ID
// @Description	Get Wardrobe
// @Tags		wardrobes
// @Accept		json
// @Produce		json
// @Param 		id		path 		string 	false 	"wardrobe id"
// @Success		200	{object}	jsonResponse{data=response.WardrobeResponse}
// @Router		/v1/wardrobe/{id}	[get]
func (api *API) GetById(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetById")
	defer span.End()

	wardrobeIDStr := req.Params("id")
	if wardrobeIDStr == "" {
		return custresp.CustomErrorResponse(errors.New("missing id"))
	}

	wardrobeID, err := uuid.Parse(wardrobeIDStr)
	if err != nil {
		return custresp.CustomErrorResponse(errors.New("invalid id"))
	}

	res, err := api.wardrobeUc.GetWardrobe(ctx, &wardrobeID)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// Delete godoc
// @Summary 	Delete Wardrobe By ID
// @Description	Delete Wardrobe
// @Tags		wardrobes
// @Accept		json
// @Produce		json
// @Param 		id		path 		string 	false 	"wardrobe id"
// @Success		200	{object}	jsonResponse{}
// @Router		/v1/wardrobe/{id}	[delete]
func (api *API) Delete(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.Delete")
	defer span.End()

	wardrobeIDStr := req.Params("id")
	if wardrobeIDStr == "" {
		return custresp.CustomErrorResponse(errors.New("missing id"))
	}

	wardrobeID, err := uuid.Parse(wardrobeIDStr)
	if err != nil {
		return custresp.CustomErrorResponse(errors.New("invalid id"))
	}

	err = api.wardrobeUc.DeleteWardrobe(ctx, &wardrobeID)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData("success"), nil
}

// Search godoc
// @Summary 	Search Wardrobe
// @Description	Search Wardrobe by color and/or size
// @Tags		wardrobes
// @Accept		json
// @Produce		json
// @Param 		color	query		string	false	"Color of the wardrobe"
// @Param 		size	query		string	false	"Size of the wardrobe"
// @Success		200		{object}	jsonResponse{data=[]response.WardrobeResponse}
// @Router		/v1/wardrobe/search	[get]
func (api *API) Search(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.Search")
	defer span.End()

	color := req.Query("color")

	size := req.Query("size")

	wardrobes, err := api.wardrobeUc.Search(ctx, color, size)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(wardrobes), nil
}

// AddStock godoc
// @Summary 	AddStock Wardrobe
// @Description	AddStock Wardrobe
// @Tags		wardrobes
// @Accept		json
// @Param		wardrobes 		body 	request.WardrobeAddSubRequest true "WardrobeAddSubRequest Payload"
// @Produce		json
// @Param 		id		path 		string 	false 	"wardrobe id"
// @Success		200	{object}	jsonResponse{data=response.WardrobeResponse}
// @Router		/v1/wardrobe/{id}/addStock	[put]
func (api *API) AddStock(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.AddStock")
	defer span.End()

	wardrobeIDStr := req.Params("id")
	if wardrobeIDStr == "" {
		return custresp.CustomErrorResponse(errors.New("missing id"))
	}

	wardrobeID, err := uuid.Parse(wardrobeIDStr)
	if err != nil {
		return custresp.CustomErrorResponse(errors.New("invalid id"))
	}

	var updateReq request.WardrobeAddSubRequest
	err = json.Unmarshal(req.RawBody(), &updateReq)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	res, err := api.wardrobeUc.AddStock(ctx, &wardrobeID, updateReq.Amount)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// SubStock godoc
// @Summary 	SubStock Wardrobe
// @Description	SubStock Wardrobe
// @Tags		wardrobes
// @Accept		json
// @Param		wardrobes 		body 	request.WardrobeAddSubRequest true "WardrobeAddSubRequest Payload"
// @Produce		json
// @Param 		id		path 		string 	false 	"wardrobe id"
// @Success		200	{object}	jsonResponse{data=response.WardrobeResponse}
// @Router		/v1/wardrobe/{id}/subStock	[put]
func (api *API) SubStock(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.AddStock")
	defer span.End()

	wardrobeIDStr := req.Params("id")
	if wardrobeIDStr == "" {
		return custresp.CustomErrorResponse(errors.New("missing id"))
	}

	wardrobeID, err := uuid.Parse(wardrobeIDStr)
	if err != nil {
		return custresp.CustomErrorResponse(errors.New("invalid id"))
	}

	var updateReq request.WardrobeAddSubRequest
	err = json.Unmarshal(req.RawBody(), &updateReq)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	res, err := api.wardrobeUc.SubStock(ctx, &wardrobeID, updateReq.Amount)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetAvailable godoc
// @Summary 	Get Available Wardrobe
// @Description	Get Available Wardrobe
// @Tags		wardrobes
// @Accept		json
// @Produce		json
// @Success		200	{object}	jsonResponse{data=response.WardrobeResponse}
// @Router		/v1/wardrobe/ready	[get]
func (api *API) GetAvailable(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetAvailable")
	defer span.End()

	res, err := api.wardrobeUc.GetAvailable(ctx)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetUnavailable godoc
// @Summary 	Get UnavailableWardrobe
// @Description	Get Unavailable Wardrobe
// @Tags		wardrobes
// @Accept		json
// @Produce		json
// @Success		200	{object}	jsonResponse{data=response.WardrobeResponse}
// @Router		/v1/wardrobe/out	[get]
func (api *API) GetUnavailable(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetAvailable")
	defer span.End()

	res, err := api.wardrobeUc.GetUnavailable(ctx)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetLessThan godoc
// @Summary 	Get LessThan Wardrobe
// @Description	Get LessThan Wardrobe
// @Tags		wardrobes
// @Accept		json
// @Produce		json
// @Param 		amount	query		string	false	"amount of the wardrobe"
// @Success		200	{object}	jsonResponse{data=response.WardrobeResponse}
// @Router		/v1/wardrobe/less	[get]
func (api *API) GetLessThan(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetAvailable")
	defer span.End()

	amount := req.Query("amount")
	amountInt, err := strconv.Atoi(amount)

	res, err := api.wardrobeUc.GetLessThan(ctx, amountInt)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}
