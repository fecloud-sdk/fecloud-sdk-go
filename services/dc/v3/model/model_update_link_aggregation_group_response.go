package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateLinkAggregationGroupResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	LinkAggregationGroup *LinkAggregationGroup `json:"link_aggregation_group,omitempty"`
	HttpStatusCode       int                   `json:"-"`
}

func (o UpdateLinkAggregationGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLinkAggregationGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateLinkAggregationGroupResponse", string(data)}, " ")
}
