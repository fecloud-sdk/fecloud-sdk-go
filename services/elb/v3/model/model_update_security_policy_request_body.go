package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// UpdateSecurityPolicyRequestBody This is a auto create Body Object
type UpdateSecurityPolicyRequestBody struct {
	SecurityPolicy *UpdateSecurityPolicyOption `json:"security_policy"`
}

func (o UpdateSecurityPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSecurityPolicyRequestBody", string(data)}, " ")
}
