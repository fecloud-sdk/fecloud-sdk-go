package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SetDatabaseUserPrivilegeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetDatabaseUserPrivilegeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetDatabaseUserPrivilegeResponse struct{}"
	}

	return strings.Join([]string{"SetDatabaseUserPrivilegeResponse", string(data)}, " ")
}
