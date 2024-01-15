package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateMemberReq struct {
	TenantId *string `json:"tenant_id,omitempty"`

	Name *string `json:"name,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	ProtocolPort int32 `json:"protocol_port"`

	SubnetId string `json:"subnet_id"`

	Address string `json:"address"`

	Weight *int32 `json:"weight,omitempty"`
}

func (o CreateMemberReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMemberReq struct{}"
	}

	return strings.Join([]string{"CreateMemberReq", string(data)}, " ")
}
