package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListLoadbalancersByTagsRequestBody struct {
	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Action string `json:"action"`

	Matches *[]ActionMatch `json:"matches,omitempty"`

	Tags *[]ActionTag `json:"tags,omitempty"`
}

func (o ListLoadbalancersByTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLoadbalancersByTagsRequestBody struct{}"
	}

	return strings.Join([]string{"ListLoadbalancersByTagsRequestBody", string(data)}, " ")
}
