package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowPrivateNatRequest struct {
	GatewayId string `json:"gateway_id"`
}

func (o ShowPrivateNatRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivateNatRequest struct{}"
	}

	return strings.Join([]string{"ShowPrivateNatRequest", string(data)}, " ")
}
