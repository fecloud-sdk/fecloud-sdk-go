package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/sdktime"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type TransitIp struct {
	Id string `json:"id"`

	ProjectId string `json:"project_id"`

	NetworkInterfaceId string `json:"network_interface_id"`

	IpAddress string `json:"ip_address"`

	CreatedAt *sdktime.SdkTime `json:"created_at"`

	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	VirsubnetId *string `json:"virsubnet_id,omitempty"`

	Tags *[]PrivateTag `json:"tags,omitempty"`

	GatewayId string `json:"gateway_id"`

	EnterpriseProjectId string `json:"enterprise_project_id"`
}

func (o TransitIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransitIp struct{}"
	}

	return strings.Join([]string{"TransitIp", string(data)}, " ")
}
