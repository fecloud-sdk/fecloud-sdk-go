package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type MemberResp struct {
	Id string `json:"id"`

	ProjectId string `json:"project_id"`

	TenantId string `json:"tenant_id"`

	Name string `json:"name"`

	AdminStateUp bool `json:"admin_state_up"`

	ProtocolPort int32 `json:"protocol_port"`

	SubnetId string `json:"subnet_id"`

	Address string `json:"address"`

	Weight int32 `json:"weight"`

	OperatingStatus string `json:"operating_status"`
}

func (o MemberResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberResp struct{}"
	}

	return strings.Join([]string{"MemberResp", string(data)}, " ")
}
