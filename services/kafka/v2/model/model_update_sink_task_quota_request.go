package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateSinkTaskQuotaRequest struct {
	ConnectorId string `json:"connector_id"`

	Body *UpdateSinkTaskQuotaReq `json:"body,omitempty"`
}

func (o UpdateSinkTaskQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSinkTaskQuotaRequest struct{}"
	}

	return strings.Join([]string{"UpdateSinkTaskQuotaRequest", string(data)}, " ")
}
