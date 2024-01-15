package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowInstanceTopicDetailRespPartitions struct {
	Partition *int32 `json:"partition,omitempty"`

	Leader *int32 `json:"leader,omitempty"`

	Leo *int32 `json:"leo,omitempty"`

	Hw *int32 `json:"hw,omitempty"`

	Lso *int32 `json:"lso,omitempty"`

	LastUpdateTimestamp *int64 `json:"last_update_timestamp,omitempty"`

	Replicas *[]ShowInstanceTopicDetailRespReplicas `json:"replicas,omitempty"`
}

func (o ShowInstanceTopicDetailRespPartitions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceTopicDetailRespPartitions struct{}"
	}

	return strings.Join([]string{"ShowInstanceTopicDetailRespPartitions", string(data)}, " ")
}
