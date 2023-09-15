package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteL7RuleResponse Response Object
type DeleteL7RuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteL7RuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteL7RuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteL7RuleResponse", string(data)}, " ")
}
