package chargepoint

import (
	"github.com/pkg/errors"
)

type Validator interface {
	validate() error
}

// createChargePointRequest
type createChargePointRequest struct {
	ChargePointSerialNumber *string `json:"chargePointSerialNumber"`
	CreateTime              *string `json:"CreateTime"`
}

func (request *createChargePointRequest) validate() error {
	number := request.ChargePointSerialNumber
	heartTime := request.CreateTime

	if number == nil {
		return errors.New("ChargePointSerialNumber 不能为空")
	}
	if heartTime == nil {
		return errors.New("CreateTime 不能为空")
	}
	return nil
}

// deleteChargePointRequest
type deleteChargePointRequest struct {
	ChargePointSerialNumber *string `json:"chargePointSerialNumber"`
}

func (request *deleteChargePointRequest) validate() error {
	number := request.ChargePointSerialNumber

	if number == nil {
		return errors.New("ChargePointSerialNumber 不能为空")
	}

	return nil
}
