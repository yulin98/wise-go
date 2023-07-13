package chargepoint

import (
	"com.wisecharge/central/internal/package/core"
	"com.wisecharge/central/internal/package/pojo/response"
	"net/http"
)

// CreateChargePoint createChargePoint
func (h *handler) CreateChargePoint() core.HandlerFunc {

	return func(ctx core.Context) {
		request := new(createChargePointRequest)
		if err := ctx.ShouldBindJSON(request); err != nil {
			apiResponse := response.ApiResponse{
				Code:    http.StatusOK,
				Message: err.Error(),
				Data:    nil,
			}
			ctx.Payload(apiResponse)
			return
		}
		if err := request.validate(); err != nil {
			apiResponse := response.ApiResponse{
				Code:    http.StatusOK,
				Message: err.Error(),
				Data:    nil,
			}
			ctx.Payload(apiResponse)
			return
		}

	}
}

// DeleteChargePoint deleteConnector
func (h *handler) DeleteChargePoint() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}

// UpdateChargePoint updateConnector
func (h *handler) UpdateChargePoint() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}

// QueryOneChargePoint queryOneConnector
func (h *handler) QueryOneChargePoint() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}

// QueryPageChargePoint queryPageChargePoint
func (h *handler) QueryPageChargePoint() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
