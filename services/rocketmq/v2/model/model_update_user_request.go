package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateUserRequest struct {
	InstanceId string `json:"instance_id"`

	UserName string `json:"user_name"`

	Body *User `json:"body,omitempty"`
}

func (o UpdateUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserRequest struct{}"
	}

	return strings.Join([]string{"UpdateUserRequest", string(data)}, " ")
}
