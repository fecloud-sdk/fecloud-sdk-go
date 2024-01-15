package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateHealthmonitorRequest struct {
	HealthmonitorId string `json:"healthmonitor_id"`

	Body *UpdateHealthmonitorRequestBody `json:"body,omitempty"`
}

func (o UpdateHealthmonitorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHealthmonitorRequest struct{}"
	}

	return strings.Join([]string{"UpdateHealthmonitorRequest", string(data)}, " ")
}
