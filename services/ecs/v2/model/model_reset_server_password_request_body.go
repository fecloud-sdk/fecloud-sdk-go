package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResetServerPasswordRequestBody struct {
	ResetPassword *ResetServerPasswordOption `json:"reset-password"`
}

func (o ResetServerPasswordRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetServerPasswordRequestBody struct{}"
	}

	return strings.Join([]string{"ResetServerPasswordRequestBody", string(data)}, " ")
}
