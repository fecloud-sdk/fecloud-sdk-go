package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowSinkTaskDetailRespPartitions struct {
	PartitionId *string `json:"partition_id,omitempty"`

	Status *string `json:"status,omitempty"`

	LastTransferOffset *string `json:"last_transfer_offset,omitempty"`

	LogEndOffset *string `json:"log_end_offset,omitempty"`

	Lag *string `json:"lag,omitempty"`
}

func (o ShowSinkTaskDetailRespPartitions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSinkTaskDetailRespPartitions struct{}"
	}

	return strings.Join([]string{"ShowSinkTaskDetailRespPartitions", string(data)}, " ")
}
