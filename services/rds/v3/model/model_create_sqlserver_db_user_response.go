package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateSqlserverDbUserResponse struct {
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSqlserverDbUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlserverDbUserResponse struct{}"
	}

	return strings.Join([]string{"CreateSqlserverDbUserResponse", string(data)}, " ")
}
