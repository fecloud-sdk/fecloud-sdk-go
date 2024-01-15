package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListTransitIpsRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	PageReverse *bool `json:"page_reverse,omitempty"`

	Id *[]string `json:"id,omitempty"`

	NetworkInterfaceId *[]string `json:"network_interface_id,omitempty"`

	IpAddress *[]string `json:"ip_address,omitempty"`

	GatewayId *[]string `json:"gateway_id,omitempty"`

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	VirsubnetId *[]string `json:"virsubnet_id,omitempty"`
}

func (o ListTransitIpsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransitIpsRequest struct{}"
	}

	return strings.Join([]string{"ListTransitIpsRequest", string(data)}, " ")
}
