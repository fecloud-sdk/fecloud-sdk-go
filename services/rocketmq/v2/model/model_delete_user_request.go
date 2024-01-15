package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteUserRequest struct {
	InstanceId string `json:"instance_id"`

	UserName string `json:"user_name"`
}

func (o DeleteUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteUserRequest struct{}"
	}

	return strings.Join([]string{"DeleteUserRequest", string(data)}, " ")
}
