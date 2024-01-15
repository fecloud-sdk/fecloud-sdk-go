package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateDatabaseUserRequest struct {
	InstanceId string `json:"instance_id"`

	Body *CreateDatabaseUserRequestBody `json:"body,omitempty"`
}

func (o CreateDatabaseUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseUserRequest struct{}"
	}

	return strings.Join([]string{"CreateDatabaseUserRequest", string(data)}, " ")
}
