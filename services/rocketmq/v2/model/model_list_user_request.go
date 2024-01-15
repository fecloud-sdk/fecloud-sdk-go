package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListUserRequest struct {
	InstanceId string `json:"instance_id"`

	Offset *string `json:"offset,omitempty"`

	Limit *string `json:"limit,omitempty"`
}

func (o ListUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserRequest struct{}"
	}

	return strings.Join([]string{"ListUserRequest", string(data)}, " ")
}
