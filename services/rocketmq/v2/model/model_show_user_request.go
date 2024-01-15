package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowUserRequest struct {
	InstanceId string `json:"instance_id"`

	UserName string `json:"user_name"`
}

func (o ShowUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserRequest struct{}"
	}

	return strings.Join([]string{"ShowUserRequest", string(data)}, " ")
}
