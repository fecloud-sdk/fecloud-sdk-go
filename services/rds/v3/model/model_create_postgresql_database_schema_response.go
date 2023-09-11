package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreatePostgresqlDatabaseSchemaResponse Response Object
type CreatePostgresqlDatabaseSchemaResponse struct {

	// 操作结果。
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePostgresqlDatabaseSchemaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostgresqlDatabaseSchemaResponse struct{}"
	}

	return strings.Join([]string{"CreatePostgresqlDatabaseSchemaResponse", string(data)}, " ")
}
