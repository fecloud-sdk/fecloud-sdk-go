package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateL7RuleRequestBody struct {
	Rule *CreateRuleOption `json:"rule"`
}

func (o CreateL7RuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateL7RuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateL7RuleRequestBody", string(data)}, " ")
}
