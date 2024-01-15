package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateOrDeleteKafkaTagRequest struct {
	InstanceId string `json:"instance_id"`

	Body *BatchCreateOrDeleteTagReq `json:"body,omitempty"`
}

func (o BatchCreateOrDeleteKafkaTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteKafkaTagRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteKafkaTagRequest", string(data)}, " ")
}
