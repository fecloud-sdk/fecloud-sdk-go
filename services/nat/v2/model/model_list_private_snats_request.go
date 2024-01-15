package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListPrivateSnatsRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	PageReverse *bool `json:"page_reverse,omitempty"`

	Id *[]string `json:"id,omitempty"`

	Description *[]string `json:"description,omitempty"`

	GatewayId *[]string `json:"gateway_id,omitempty"`

	Cidr *[]string `json:"cidr,omitempty"`

	VirsubnetId *[]string `json:"virsubnet_id,omitempty"`

	TransitIpId *[]string `json:"transit_ip_id,omitempty"`

	TransitIpAddress *[]string `json:"transit_ip_address,omitempty"`

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
}

func (o ListPrivateSnatsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateSnatsRequest struct{}"
	}

	return strings.Join([]string{"ListPrivateSnatsRequest", string(data)}, " ")
}
