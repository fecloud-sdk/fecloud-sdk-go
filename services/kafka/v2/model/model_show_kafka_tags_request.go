package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowKafkaTagsRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o ShowKafkaTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKafkaTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowKafkaTagsRequest", string(data)}, " ")
}
