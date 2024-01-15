package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteDatabaseUserRequest struct {
	InstanceId string `json:"instance_id"`

	Body *DeleteDatabaseUserRequestBody `json:"body,omitempty"`
}

func (o DeleteDatabaseUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabaseUserRequest struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseUserRequest", string(data)}, " ")
}
