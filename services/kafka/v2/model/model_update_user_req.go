package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateUserReq struct {
	NewPassword *string `json:"new_password,omitempty"`

	UserName *string `json:"user_name,omitempty"`

	UserDesc *string `json:"user_desc,omitempty"`
}

func (o UpdateUserReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserReq struct{}"
	}

	return strings.Join([]string{"UpdateUserReq", string(data)}, " ")
}
