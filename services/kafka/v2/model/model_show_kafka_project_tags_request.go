package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowKafkaProjectTagsRequest struct {
}

func (o ShowKafkaProjectTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKafkaProjectTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowKafkaProjectTagsRequest", string(data)}, " ")
}
