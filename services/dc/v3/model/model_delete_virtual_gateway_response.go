package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteVirtualGatewayResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVirtualGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVirtualGatewayResponse struct{}"
	}

	return strings.Join([]string{"DeleteVirtualGatewayResponse", string(data)}, " ")
}
