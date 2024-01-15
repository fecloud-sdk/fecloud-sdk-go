package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowL7ruleRequest struct {
	L7policyId string `json:"l7policy_id"`

	L7ruleId string `json:"l7rule_id"`
}

func (o ShowL7ruleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowL7ruleRequest struct{}"
	}

	return strings.Join([]string{"ShowL7ruleRequest", string(data)}, " ")
}
