package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListDatabaseUsersRequest struct {
	InstanceId string `json:"instance_id"`

	UserName *string `json:"user_name,omitempty"`

	DbName *string `json:"db_name,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDatabaseUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseUsersRequest struct{}"
	}

	return strings.Join([]string{"ListDatabaseUsersRequest", string(data)}, " ")
}
