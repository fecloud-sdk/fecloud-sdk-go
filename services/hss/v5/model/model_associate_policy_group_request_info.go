package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AssociatePolicyGroupRequestInfo struct {
	TargetPolicyGroupId string `json:"target_policy_group_id"`

	OperateAll *bool `json:"operate_all,omitempty"`

	HostIdList *[]string `json:"host_id_list,omitempty"`
}

func (o AssociatePolicyGroupRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociatePolicyGroupRequestInfo struct{}"
	}

	return strings.Join([]string{"AssociatePolicyGroupRequestInfo", string(data)}, " ")
}
