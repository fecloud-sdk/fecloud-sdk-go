package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowL7ruleResponse struct {
	Rule           *L7ruleResp `json:"rule,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowL7ruleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowL7ruleResponse struct{}"
	}

	return strings.Join([]string{"ShowL7ruleResponse", string(data)}, " ")
}
