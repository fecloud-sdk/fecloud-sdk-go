package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/sdktime"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type PrivateSnat struct {
	Id *string `json:"id,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	GatewayId *string `json:"gateway_id,omitempty"`

	Cidr *string `json:"cidr,omitempty"`

	VirsubnetId *string `json:"virsubnet_id,omitempty"`

	Description *string `json:"description,omitempty"`

	TransitIpAssociations *[]AssociatedTransitIp `json:"transit_ip_associations,omitempty"`

	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o PrivateSnat) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateSnat struct{}"
	}

	return strings.Join([]string{"PrivateSnat", string(data)}, " ")
}
