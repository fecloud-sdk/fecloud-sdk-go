package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteSecurityPolicyRequest Request Object
type DeleteSecurityPolicyRequest struct {

	// 自定义安全策略的ID。
	SecurityPolicyId string `json:"security_policy_id"`
}

func (o DeleteSecurityPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityPolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteSecurityPolicyRequest", string(data)}, " ")
}
