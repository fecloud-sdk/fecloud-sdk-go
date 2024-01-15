package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResetPasswordReq struct {
	NewPassword string `json:"new_password"`
}

func (o ResetPasswordReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetPasswordReq struct{}"
	}

	return strings.Join([]string{"ResetPasswordReq", string(data)}, " ")
}
