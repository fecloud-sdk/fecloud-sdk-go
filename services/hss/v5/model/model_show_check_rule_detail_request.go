package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowCheckRuleDetailRequest struct {
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	CheckName string `json:"check_name"`

	CheckType string `json:"check_type"`

	CheckRuleId string `json:"check_rule_id"`

	Standard string `json:"standard"`

	HostId *string `json:"host_id,omitempty"`
}

func (o ShowCheckRuleDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCheckRuleDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowCheckRuleDetailRequest", string(data)}, " ")
}
