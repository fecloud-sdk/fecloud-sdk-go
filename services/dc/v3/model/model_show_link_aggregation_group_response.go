package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowLinkAggregationGroupResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	LinkAggregationGroup *LinkAggregationGroup `json:"link_aggregation_group,omitempty"`
	HttpStatusCode       int                   `json:"-"`
}

func (o ShowLinkAggregationGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLinkAggregationGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowLinkAggregationGroupResponse", string(data)}, " ")
}
