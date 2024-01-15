package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateInstanceConsumerGroupRequest struct {
	Engine string `json:"engine"`

	InstanceId string `json:"instance_id"`

	Group string `json:"group"`

	Body *CreateGroupReq `json:"body,omitempty"`
}

func (o UpdateInstanceConsumerGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceConsumerGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceConsumerGroupRequest", string(data)}, " ")
}
