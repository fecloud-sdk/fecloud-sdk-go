package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteDatabaseRoleRequest struct {
	InstanceId string `json:"instance_id"`

	Body *DeleteDatabaseRoleRequestBody `json:"body,omitempty"`
}

func (o DeleteDatabaseRoleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabaseRoleRequest struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseRoleRequest", string(data)}, " ")
}
