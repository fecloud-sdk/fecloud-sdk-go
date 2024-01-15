package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteInstanceUsersResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteInstanceUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInstanceUsersResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstanceUsersResponse", string(data)}, " ")
}
