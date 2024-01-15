package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DbUserPrivilegeRequest struct {
	UserName string `json:"user_name"`

	AuthorizationType string `json:"authorization_type"`

	Privileges []string `json:"privileges"`
}

func (o DbUserPrivilegeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbUserPrivilegeRequest struct{}"
	}

	return strings.Join([]string{"DbUserPrivilegeRequest", string(data)}, " ")
}
