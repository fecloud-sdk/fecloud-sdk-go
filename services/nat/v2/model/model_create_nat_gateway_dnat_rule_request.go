package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateNatGatewayDnatRuleRequest struct {
	Body *CreateNatGatewayDnatRuleOption `json:"body,omitempty"`
}

func (o CreateNatGatewayDnatRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNatGatewayDnatRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateNatGatewayDnatRuleRequest", string(data)}, " ")
}
