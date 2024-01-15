package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowRiskConfigDetailResponse struct {
	Severity *string `json:"severity,omitempty"`

	CheckType *string `json:"check_type,omitempty"`

	CheckTypeDesc *string `json:"check_type_desc,omitempty"`

	CheckRuleNum *int32 `json:"check_rule_num,omitempty"`

	FailedRuleNum *int32 `json:"failed_rule_num,omitempty"`

	PassedRuleNum *int32 `json:"passed_rule_num,omitempty"`

	IgnoredRuleNum *int32 `json:"ignored_rule_num,omitempty"`

	HostNum        *int64 `json:"host_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowRiskConfigDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRiskConfigDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowRiskConfigDetailResponse", string(data)}, " ")
}
