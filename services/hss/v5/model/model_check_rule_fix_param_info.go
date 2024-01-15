package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CheckRuleFixParamInfo struct {
	RuleParamId *int32 `json:"rule_param_id,omitempty"`

	RuleDesc *string `json:"rule_desc,omitempty"`

	DefaultValue *int32 `json:"default_value,omitempty"`

	RangeMin *int32 `json:"range_min,omitempty"`

	RangeMax *int32 `json:"range_max,omitempty"`
}

func (o CheckRuleFixParamInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRuleFixParamInfo struct{}"
	}

	return strings.Join([]string{"CheckRuleFixParamInfo", string(data)}, " ")
}
