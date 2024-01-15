package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowL7policyRequest struct {
	L7policyId string `json:"l7policy_id"`
}

func (o ShowL7policyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowL7policyRequest struct{}"
	}

	return strings.Join([]string{"ShowL7policyRequest", string(data)}, " ")
}
