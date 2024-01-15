package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreatTransitIpOption struct {
	VirsubnetId string `json:"virsubnet_id"`

	IpAddress *string `json:"ip_address,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Tags *[]PrivateTag `json:"tags,omitempty"`
}

func (o CreatTransitIpOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatTransitIpOption struct{}"
	}

	return strings.Join([]string{"CreatTransitIpOption", string(data)}, " ")
}
