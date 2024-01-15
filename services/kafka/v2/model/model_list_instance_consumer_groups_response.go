package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListInstanceConsumerGroupsResponse struct {
	Groups *[]GroupInfoSimple `json:"groups,omitempty"`

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceConsumerGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceConsumerGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceConsumerGroupsResponse", string(data)}, " ")
}
