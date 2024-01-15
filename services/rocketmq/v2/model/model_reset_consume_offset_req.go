package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResetConsumeOffsetReq struct {
	Topic *string `json:"topic,omitempty"`

	Timestamp float32 `json:"timestamp,omitempty"`
}

func (o ResetConsumeOffsetReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetConsumeOffsetReq struct{}"
	}

	return strings.Join([]string{"ResetConsumeOffsetReq", string(data)}, " ")
}
