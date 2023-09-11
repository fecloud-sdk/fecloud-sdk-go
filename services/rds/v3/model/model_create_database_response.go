package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateDatabaseResponse Response Object
type CreateDatabaseResponse struct {

	// 操作结果。
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseResponse struct{}"
	}

	return strings.Join([]string{"CreateDatabaseResponse", string(data)}, " ")
}
