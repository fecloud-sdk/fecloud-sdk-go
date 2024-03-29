package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateDatabaseUserResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateDatabaseUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseUserResponse struct{}"
	}

	return strings.Join([]string{"CreateDatabaseUserResponse", string(data)}, " ")
}
