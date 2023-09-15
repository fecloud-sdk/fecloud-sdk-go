package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// RuleRef
type RuleRef struct {

	// 规则ID。
	Id string `json:"id"`
}

func (o RuleRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleRef struct{}"
	}

	return strings.Join([]string{"RuleRef", string(data)}, " ")
}
