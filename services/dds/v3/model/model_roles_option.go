package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RolesOption struct {
	RoleDbName string `json:"role_db_name"`

	RoleName string `json:"role_name"`
}

func (o RolesOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RolesOption struct{}"
	}

	return strings.Join([]string{"RolesOption", string(data)}, " ")
}
