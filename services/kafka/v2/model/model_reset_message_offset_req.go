package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResetMessageOffsetReq struct {
	Topic string `json:"topic"`

	Partition *int32 `json:"partition,omitempty"`

	MessageOffset *int32 `json:"message_offset,omitempty"`

	Timestamp *int32 `json:"timestamp,omitempty"`
}

func (o ResetMessageOffsetReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetMessageOffsetReq struct{}"
	}

	return strings.Join([]string{"ResetMessageOffsetReq", string(data)}, " ")
}
