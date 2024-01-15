package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateInstanceUserReq struct {
	UserName *string `json:"user_name,omitempty"`

	UserDesc *string `json:"user_desc,omitempty"`

	UserPasswd *string `json:"user_passwd,omitempty"`
}

func (o CreateInstanceUserReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceUserReq struct{}"
	}

	return strings.Join([]string{"CreateInstanceUserReq", string(data)}, " ")
}
