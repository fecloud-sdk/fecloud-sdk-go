package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResetConsumeOffsetRespQueues struct {
	BrokerName *string `json:"broker_name,omitempty"`

	QueueId *int32 `json:"queue_id,omitempty"`

	TimestampOffset *int64 `json:"timestamp_offset,omitempty"`
}

func (o ResetConsumeOffsetRespQueues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetConsumeOffsetRespQueues struct{}"
	}

	return strings.Join([]string{"ResetConsumeOffsetRespQueues", string(data)}, " ")
}
