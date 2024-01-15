package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListVirtualGatewaysResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	VirtualGateways *[]VirtualGateway `json:"virtual_gateways,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListVirtualGatewaysResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVirtualGatewaysResponse struct{}"
	}

	return strings.Join([]string{"ListVirtualGatewaysResponse", string(data)}, " ")
}
