package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateL7RuleRequestBody struct {
	Rule *UpdateL7RuleOption `json:"rule"`
}

func (o UpdateL7RuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateL7RuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateL7RuleRequestBody", string(data)}, " ")
}
