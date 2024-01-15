package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeletePrivateNatRequest struct {
	GatewayId string `json:"gateway_id"`
}

func (o DeletePrivateNatRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateNatRequest struct{}"
	}

	return strings.Join([]string{"DeletePrivateNatRequest", string(data)}, " ")
}
