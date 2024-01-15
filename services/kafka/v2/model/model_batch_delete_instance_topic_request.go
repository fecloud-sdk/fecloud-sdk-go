package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteInstanceTopicRequest struct {
	InstanceId string `json:"instance_id"`

	Body *BatchDeleteInstanceTopicReq `json:"body,omitempty"`
}

func (o BatchDeleteInstanceTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInstanceTopicRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstanceTopicRequest", string(data)}, " ")
}
