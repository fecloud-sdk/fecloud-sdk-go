package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DbUserPwdRequest struct {
	Name string `json:"name"`

	Password string `json:"password"`
}

func (o DbUserPwdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbUserPwdRequest struct{}"
	}

	return strings.Join([]string{"DbUserPwdRequest", string(data)}, " ")
}
