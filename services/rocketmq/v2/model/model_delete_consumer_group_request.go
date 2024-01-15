package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteConsumerGroupRequest struct {
	InstanceId string `json:"instance_id"`

	Group string `json:"group"`
}

func (o DeleteConsumerGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConsumerGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteConsumerGroupRequest", string(data)}, " ")
}
