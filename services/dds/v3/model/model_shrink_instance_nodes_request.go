package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShrinkInstanceNodesRequest struct {
	InstanceId string `json:"instance_id"`

	Body *ReduceInstanceNodeRequestBody `json:"body,omitempty"`
}

func (o ShrinkInstanceNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShrinkInstanceNodesRequest struct{}"
	}

	return strings.Join([]string{"ShrinkInstanceNodesRequest", string(data)}, " ")
}
