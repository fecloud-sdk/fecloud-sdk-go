package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// NovaSecurityGroupCommonGroup
type NovaSecurityGroupCommonGroup struct {

	// 对端安全组的名称
	Name *string `json:"name,omitempty"`

	// 对端安全组所属租户的租户ID
	TenantId *string `json:"tenant_id,omitempty"`
}

func (o NovaSecurityGroupCommonGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaSecurityGroupCommonGroup struct{}"
	}

	return strings.Join([]string{"NovaSecurityGroupCommonGroup", string(data)}, " ")
}
