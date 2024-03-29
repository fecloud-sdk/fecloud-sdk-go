package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateL7ruleRequestBody struct {
	Rule *UpdateL7ruleReq `json:"rule"`
}

func (o UpdateL7ruleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateL7ruleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateL7ruleRequestBody", string(data)}, " ")
}
