package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateDatabaseUserRequestBody struct {
	UserName string `json:"user_name"`

	UserPwd string `json:"user_pwd"`

	DbName *string `json:"db_name,omitempty"`

	Roles []RolesOption `json:"roles"`
}

func (o CreateDatabaseUserRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseUserRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDatabaseUserRequestBody", string(data)}, " ")
}
