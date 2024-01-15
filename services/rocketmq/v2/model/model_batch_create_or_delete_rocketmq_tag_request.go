package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateOrDeleteRocketmqTagRequest struct {
	InstanceId string `json:"instance_id"`

	Body *BatchCreateOrDeleteTagReq `json:"body,omitempty"`
}

func (o BatchCreateOrDeleteRocketmqTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteRocketmqTagRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteRocketmqTagRequest", string(data)}, " ")
}
