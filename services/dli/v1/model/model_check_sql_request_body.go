package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CheckSqlRequestBody struct {
	Sql string `json:"sql"`

	Currentdb *string `json:"currentdb,omitempty"`
}

func (o CheckSqlRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckSqlRequestBody struct{}"
	}

	return strings.Join([]string{"CheckSqlRequestBody", string(data)}, " ")
}
