package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateL7ruleResponse struct {
	Rule           *L7ruleResp `json:"rule,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o CreateL7ruleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateL7ruleResponse struct{}"
	}

	return strings.Join([]string{"CreateL7ruleResponse", string(data)}, " ")
}
