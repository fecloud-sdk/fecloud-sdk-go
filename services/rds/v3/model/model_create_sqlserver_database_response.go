package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateSqlserverDatabaseResponse struct {
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSqlserverDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlserverDatabaseResponse struct{}"
	}

	return strings.Join([]string{"CreateSqlserverDatabaseResponse", string(data)}, " ")
}
