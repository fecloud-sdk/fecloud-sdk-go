package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListPrivateDnatsRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	PageReverse *bool `json:"page_reverse,omitempty"`

	Id *[]string `json:"id,omitempty"`

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	Description *[]string `json:"description,omitempty"`

	GatewayId *[]string `json:"gateway_id,omitempty"`

	TransitIpId *[]string `json:"transit_ip_id,omitempty"`

	ExternalIpAddress *[]string `json:"external_ip_address,omitempty"`

	NetworkInterfaceId *[]string `json:"network_interface_id,omitempty"`

	Type *[]string `json:"type,omitempty"`

	PrivateIpAddress *[]string `json:"private_ip_address,omitempty"`
}

func (o ListPrivateDnatsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateDnatsRequest struct{}"
	}

	return strings.Join([]string{"ListPrivateDnatsRequest", string(data)}, " ")
}
