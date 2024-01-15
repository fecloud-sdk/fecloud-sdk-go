package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListDatabaseRolesRequest struct {
	InstanceId string `json:"instance_id"`

	RoleName *string `json:"role_name,omitempty"`

	DbName *string `json:"db_name,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDatabaseRolesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseRolesRequest struct{}"
	}

	return strings.Join([]string{"ListDatabaseRolesRequest", string(data)}, " ")
}
