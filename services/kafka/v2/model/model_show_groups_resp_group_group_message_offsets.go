package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowGroupsRespGroupGroupMessageOffsets struct {
	Partition *int32 `json:"partition,omitempty"`

	Lag *int32 `json:"lag,omitempty"`

	Topic *string `json:"topic,omitempty"`

	MessageCurrentOffset *int32 `json:"message_current_offset,omitempty"`

	MessageLogEndOffset *int32 `json:"message_log_end_offset,omitempty"`
}

func (o ShowGroupsRespGroupGroupMessageOffsets) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupsRespGroupGroupMessageOffsets struct{}"
	}

	return strings.Join([]string{"ShowGroupsRespGroupGroupMessageOffsets", string(data)}, " ")
}
