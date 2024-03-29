package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateL7ruleRequestBody struct {
	Rule *CreateL7ruleReq `json:"rule"`
}

func (o CreateL7ruleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateL7ruleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateL7ruleRequestBody", string(data)}, " ")
}
