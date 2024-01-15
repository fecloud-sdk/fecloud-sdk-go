package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateDatabaseRoleRequestBody struct {
	RoleName string `json:"role_name"`

	DbName *string `json:"db_name,omitempty"`

	Roles *[]RolesOption `json:"roles,omitempty"`
}

func (o CreateDatabaseRoleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseRoleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDatabaseRoleRequestBody", string(data)}, " ")
}
