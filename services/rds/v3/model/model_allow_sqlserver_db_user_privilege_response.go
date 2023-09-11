package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// AllowSqlserverDbUserPrivilegeResponse Response Object
type AllowSqlserverDbUserPrivilegeResponse struct {

	// 操作结果。
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AllowSqlserverDbUserPrivilegeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllowSqlserverDbUserPrivilegeResponse struct{}"
	}

	return strings.Join([]string{"AllowSqlserverDbUserPrivilegeResponse", string(data)}, " ")
}
