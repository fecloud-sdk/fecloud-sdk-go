package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResetManagerPasswordReq struct {
	NewPassword *string `json:"new_password,omitempty"`
}

func (o ResetManagerPasswordReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetManagerPasswordReq struct{}"
	}

	return strings.Join([]string{"ResetManagerPasswordReq", string(data)}, " ")
}
