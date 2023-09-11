package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteSqlserverDatabaseResponse Response Object
type DeleteSqlserverDatabaseResponse struct {

	// 操作结果。
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteSqlserverDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSqlserverDatabaseResponse struct{}"
	}

	return strings.Join([]string{"DeleteSqlserverDatabaseResponse", string(data)}, " ")
}
