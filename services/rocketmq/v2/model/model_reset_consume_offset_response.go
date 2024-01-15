package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResetConsumeOffsetResponse struct {
	Queues         *[]ResetConsumeOffsetRespQueues `json:"queues,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ResetConsumeOffsetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetConsumeOffsetResponse struct{}"
	}

	return strings.Join([]string{"ResetConsumeOffsetResponse", string(data)}, " ")
}
