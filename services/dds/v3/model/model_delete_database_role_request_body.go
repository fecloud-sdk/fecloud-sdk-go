package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteDatabaseRoleRequestBody struct {
	RoleName string `json:"role_name"`

	DbName string `json:"db_name"`
}

func (o DeleteDatabaseRoleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabaseRoleRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseRoleRequestBody", string(data)}, " ")
}
