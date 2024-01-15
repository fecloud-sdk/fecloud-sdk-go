package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowRocketmqTagsRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o ShowRocketmqTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRocketmqTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowRocketmqTagsRequest", string(data)}, " ")
}
