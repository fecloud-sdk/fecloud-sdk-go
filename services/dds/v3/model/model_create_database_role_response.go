package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateDatabaseRoleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateDatabaseRoleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseRoleResponse struct{}"
	}

	return strings.Join([]string{"CreateDatabaseRoleResponse", string(data)}, " ")
}
