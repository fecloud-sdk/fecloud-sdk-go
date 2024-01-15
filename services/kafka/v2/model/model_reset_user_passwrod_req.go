package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResetUserPasswrodReq struct {
	NewPassword *string `json:"new_password,omitempty"`
}

func (o ResetUserPasswrodReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetUserPasswrodReq struct{}"
	}

	return strings.Join([]string{"ResetUserPasswrodReq", string(data)}, " ")
}
