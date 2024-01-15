package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListLoadbalancersByTagsRequest struct {
	Body *ListLoadbalancersByTagsRequestBody `json:"body,omitempty"`
}

func (o ListLoadbalancersByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLoadbalancersByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListLoadbalancersByTagsRequest", string(data)}, " ")
}
