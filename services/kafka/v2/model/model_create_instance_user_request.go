package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateInstanceUserRequest struct {
	InstanceId string `json:"instance_id"`

	Body *CreateInstanceUserReq `json:"body,omitempty"`
}

func (o CreateInstanceUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceUserRequest struct{}"
	}

	return strings.Join([]string{"CreateInstanceUserRequest", string(data)}, " ")
}
