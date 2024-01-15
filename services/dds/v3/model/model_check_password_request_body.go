package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CheckPasswordRequestBody struct {
	UserPwd string `json:"user_pwd"`

	UserName *string `json:"user_name,omitempty"`

	DbName *string `json:"db_name,omitempty"`
}

func (o CheckPasswordRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckPasswordRequestBody struct{}"
	}

	return strings.Join([]string{"CheckPasswordRequestBody", string(data)}, " ")
}
