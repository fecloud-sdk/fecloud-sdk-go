package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowPartitionBeginningMessageResponse struct {
	Topic *string `json:"topic,omitempty"`

	Partition *int32 `json:"partition,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Timestamp      *int64 `json:"timestamp,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowPartitionBeginningMessageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPartitionBeginningMessageResponse struct{}"
	}

	return strings.Join([]string{"ShowPartitionBeginningMessageResponse", string(data)}, " ")
}
