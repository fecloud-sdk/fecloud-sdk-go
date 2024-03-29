package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResetPasswordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetPasswordResponse struct{}"
	}

	return strings.Join([]string{"ResetPasswordResponse", string(data)}, " ")
}
