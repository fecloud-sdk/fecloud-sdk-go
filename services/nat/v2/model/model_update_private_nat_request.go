package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdatePrivateNatRequest struct {
	GatewayId string `json:"gateway_id"`

	Body *UpdatePrivateNatRequestBody `json:"body,omitempty"`
}

func (o UpdatePrivateNatRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateNatRequest struct{}"
	}

	return strings.Join([]string{"UpdatePrivateNatRequest", string(data)}, " ")
}
