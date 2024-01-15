package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CheckRuleRiskInfoResponseInfo struct {
	Severity *string `json:"severity,omitempty"`

	CheckName *string `json:"check_name,omitempty"`

	CheckType *string `json:"check_type,omitempty"`

	Standard *string `json:"standard,omitempty"`

	CheckRuleName *string `json:"check_rule_name,omitempty"`

	CheckRuleId *string `json:"check_rule_id,omitempty"`

	HostNum *int32 `json:"host_num,omitempty"`

	ScanResult *string `json:"scan_result,omitempty"`

	Status *string `json:"status,omitempty"`

	EnableFix *int32 `json:"enable_fix,omitempty"`

	EnableClick *bool `json:"enable_click,omitempty"`

	RuleParams *[]CheckRuleFixParamInfo `json:"rule_params,omitempty"`
}

func (o CheckRuleRiskInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRuleRiskInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"CheckRuleRiskInfoResponseInfo", string(data)}, " ")
}
