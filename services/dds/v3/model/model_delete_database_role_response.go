package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteDatabaseRoleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDatabaseRoleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabaseRoleResponse struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseRoleResponse", string(data)}, " ")
}
