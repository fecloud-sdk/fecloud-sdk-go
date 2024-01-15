package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RuleCondition struct {
	Key string `json:"key"`

	Value string `json:"value"`
}

func (o RuleCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleCondition struct{}"
	}

	return strings.Join([]string{"RuleCondition", string(data)}, " ")
}
