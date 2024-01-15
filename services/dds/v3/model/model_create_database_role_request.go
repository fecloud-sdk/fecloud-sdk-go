package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateDatabaseRoleRequest struct {
	InstanceId string `json:"instance_id"`

	Body *CreateDatabaseRoleRequestBody `json:"body,omitempty"`
}

func (o CreateDatabaseRoleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseRoleRequest struct{}"
	}

	return strings.Join([]string{"CreateDatabaseRoleRequest", string(data)}, " ")
}
