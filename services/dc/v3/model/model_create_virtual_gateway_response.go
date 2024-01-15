package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateVirtualGatewayResponse struct {
	VirtualGateway *VirtualGateway `json:"virtual_gateway,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateVirtualGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVirtualGatewayResponse struct{}"
	}

	return strings.Join([]string{"CreateVirtualGatewayResponse", string(data)}, " ")
}
