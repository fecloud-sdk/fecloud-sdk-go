package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowVirtualGatewayResponse struct {
	VirtualGateway *VirtualGateway `json:"virtual_gateway,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowVirtualGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVirtualGatewayResponse struct{}"
	}

	return strings.Join([]string{"ShowVirtualGatewayResponse", string(data)}, " ")
}
