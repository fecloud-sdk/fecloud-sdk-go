package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type PartitionReassignRequest struct {
	Reassignments []PartitionReassignEntity `json:"reassignments"`

	Throttle *int32 `json:"throttle,omitempty"`
}

func (o PartitionReassignRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PartitionReassignRequest struct{}"
	}

	return strings.Join([]string{"PartitionReassignRequest", string(data)}, " ")
}
