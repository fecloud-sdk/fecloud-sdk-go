package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListInstanceConsumerGroupsResponse struct {
	Total float32 `json:"total,omitempty"`

	Groups *[]ConsumerGroup `json:"groups,omitempty"`

	Max float32 `json:"max,omitempty"`

	Remaining float32 `json:"remaining,omitempty"`

	NextOffset float32 `json:"next_offset,omitempty"`

	PreviousOffset float32 `json:"previous_offset,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListInstanceConsumerGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceConsumerGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceConsumerGroupsResponse", string(data)}, " ")
}
