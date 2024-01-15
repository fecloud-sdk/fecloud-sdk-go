package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListDatabaseRolesResponse struct {
	Roles *string `json:"roles,omitempty"`

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDatabaseRolesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseRolesResponse struct{}"
	}

	return strings.Join([]string{"ListDatabaseRolesResponse", string(data)}, " ")
}
