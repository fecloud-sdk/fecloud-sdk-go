package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListL7rulesResponse struct {
	Rules          *[]L7ruleResp `json:"rules,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListL7rulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListL7rulesResponse struct{}"
	}

	return strings.Join([]string{"ListL7rulesResponse", string(data)}, " ")
}
