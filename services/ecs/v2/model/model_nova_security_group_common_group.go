package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NovaSecurityGroupCommonGroup struct {
	Name *string `json:"name,omitempty"`

	TenantId *string `json:"tenant_id,omitempty"`
}

func (o NovaSecurityGroupCommonGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaSecurityGroupCommonGroup struct{}"
	}

	return strings.Join([]string{"NovaSecurityGroupCommonGroup", string(data)}, " ")
}
