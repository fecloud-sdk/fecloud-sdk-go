package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateInstanceUserRequest struct {
	Engine string `json:"engine"`

	InstanceId string `json:"instance_id"`

	UserName string `json:"user_name"`

	Body *UpdateUserReq `json:"body,omitempty"`
}

func (o UpdateInstanceUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceUserRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceUserRequest", string(data)}, " ")
}
