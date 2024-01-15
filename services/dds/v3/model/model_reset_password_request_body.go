package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResetPasswordRequestBody struct {
	UserPwd string `json:"user_pwd"`

	UserName *string `json:"user_name,omitempty"`

	DbName *string `json:"db_name,omitempty"`
}

func (o ResetPasswordRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetPasswordRequestBody struct{}"
	}

	return strings.Join([]string{"ResetPasswordRequestBody", string(data)}, " ")
}
