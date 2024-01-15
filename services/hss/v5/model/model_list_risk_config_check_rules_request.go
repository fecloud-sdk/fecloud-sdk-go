package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListRiskConfigCheckRulesRequest struct {
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	CheckName string `json:"check_name"`

	Standard string `json:"standard"`

	ResultType *string `json:"result_type,omitempty"`

	CheckRuleName *string `json:"check_rule_name,omitempty"`

	Severity *string `json:"severity,omitempty"`

	HostId *string `json:"host_id,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListRiskConfigCheckRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRiskConfigCheckRulesRequest struct{}"
	}

	return strings.Join([]string{"ListRiskConfigCheckRulesRequest", string(data)}, " ")
}
