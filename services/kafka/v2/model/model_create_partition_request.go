package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreatePartitionRequest struct {
	InstanceId string `json:"instance_id"`

	Topic string `json:"topic"`

	Body *CreatePartitionReq `json:"body,omitempty"`
}

func (o CreatePartitionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePartitionRequest struct{}"
	}

	return strings.Join([]string{"CreatePartitionRequest", string(data)}, " ")
}
