package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CheckRuleCheckCaseResponseInfo struct {
	CheckDescription *string `json:"check_description,omitempty"`

	CurrentValue *string `json:"current_value,omitempty"`

	SuggestValue *string `json:"suggest_value,omitempty"`
}

func (o CheckRuleCheckCaseResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRuleCheckCaseResponseInfo struct{}"
	}

	return strings.Join([]string{"CheckRuleCheckCaseResponseInfo", string(data)}, " ")
}
