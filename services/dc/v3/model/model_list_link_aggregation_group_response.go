package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListLinkAggregationGroupResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	LinkAggregationGroups *[]LinkAggregationGroup `json:"link_aggregation_groups,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListLinkAggregationGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLinkAggregationGroupResponse struct{}"
	}

	return strings.Join([]string{"ListLinkAggregationGroupResponse", string(data)}, " ")
}
