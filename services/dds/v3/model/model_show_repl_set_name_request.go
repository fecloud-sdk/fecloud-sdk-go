package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowReplSetNameRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o ShowReplSetNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReplSetNameRequest struct{}"
	}

	return strings.Join([]string{"ShowReplSetNameRequest", string(data)}, " ")
}
