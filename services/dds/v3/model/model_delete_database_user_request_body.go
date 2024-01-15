package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteDatabaseUserRequestBody struct {
	UserName string `json:"user_name"`

	DbName string `json:"db_name"`
}

func (o DeleteDatabaseUserRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabaseUserRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseUserRequestBody", string(data)}, " ")
}
