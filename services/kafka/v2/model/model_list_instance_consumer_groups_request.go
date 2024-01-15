package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListInstanceConsumerGroupsRequest struct {
	InstanceId string `json:"instance_id"`

	Offset *string `json:"offset,omitempty"`

	Limit *string `json:"limit,omitempty"`

	Group *string `json:"group,omitempty"`
}

func (o ListInstanceConsumerGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceConsumerGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceConsumerGroupsRequest", string(data)}, " ")
}
