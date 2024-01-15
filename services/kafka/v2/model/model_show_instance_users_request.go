package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowInstanceUsersRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o ShowInstanceUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceUsersRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceUsersRequest", string(data)}, " ")
}
