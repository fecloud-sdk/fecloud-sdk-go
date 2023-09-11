package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// UpdateNatGatewaySnatRuleResponse Response Object
type UpdateNatGatewaySnatRuleResponse struct {
	SnatRule       *NatGatewayUpdateSnatRuleResponseBody `json:"snat_rule,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o UpdateNatGatewaySnatRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatGatewaySnatRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateNatGatewaySnatRuleResponse", string(data)}, " ")
}
