package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListProtectionPolicyRequest struct {
	Region string `json:"region"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	PolicyName *string `json:"policy_name,omitempty"`

	ProtectPolicyId *string `json:"protect_policy_id,omitempty"`

	OperatingSystem *string `json:"operating_system,omitempty"`
}

func (o ListProtectionPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectionPolicyRequest struct{}"
	}

	return strings.Join([]string{"ListProtectionPolicyRequest", string(data)}, " ")
}
