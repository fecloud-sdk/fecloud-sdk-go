package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResetReplicaReqPartitions struct {
	Partition *int32 `json:"partition,omitempty"`

	Replicas *[]int32 `json:"replicas,omitempty"`
}

func (o ResetReplicaReqPartitions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetReplicaReqPartitions struct{}"
	}

	return strings.Join([]string{"ResetReplicaReqPartitions", string(data)}, " ")
}
