package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteL7ruleRequest struct {
	L7policyId string `json:"l7policy_id"`

	L7ruleId string `json:"l7rule_id"`
}

func (o DeleteL7ruleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteL7ruleRequest struct{}"
	}

	return strings.Join([]string{"DeleteL7ruleRequest", string(data)}, " ")
}
