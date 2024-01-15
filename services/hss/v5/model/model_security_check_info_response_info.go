package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SecurityCheckInfoResponseInfo struct {
	Severity *string `json:"severity,omitempty"`

	CheckName *string `json:"check_name,omitempty"`

	CheckType *string `json:"check_type,omitempty"`

	Standard *string `json:"standard,omitempty"`

	CheckRuleNum *int32 `json:"check_rule_num,omitempty"`

	FailedRuleNum *int32 `json:"failed_rule_num,omitempty"`

	HostNum *int32 `json:"host_num,omitempty"`

	ScanTime *int64 `json:"scan_time,omitempty"`

	CheckTypeDesc *string `json:"check_type_desc,omitempty"`
}

func (o SecurityCheckInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityCheckInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"SecurityCheckInfoResponseInfo", string(data)}, " ")
}
