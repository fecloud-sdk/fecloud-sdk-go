package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateInstanceReq struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	SecurityGroupId *string `json:"security_group_id,omitempty"`

	EnableAcl *bool `json:"enable_acl,omitempty"`

	EnablePublicip *bool `json:"enable_publicip,omitempty"`

	PublicipId *string `json:"publicip_id,omitempty"`
}

func (o UpdateInstanceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceReq struct{}"
	}

	return strings.Join([]string{"UpdateInstanceReq", string(data)}, " ")
}
