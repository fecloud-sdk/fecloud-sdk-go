package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ResetPwdResponse Response Object
type ResetPwdResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetPwdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetPwdResponse struct{}"
	}

	return strings.Join([]string{"ResetPwdResponse", string(data)}, " ")
}
