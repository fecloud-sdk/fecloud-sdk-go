package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateLinkAggregationGroupResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	LinkAggregationGroup *LinkAggregationGroup `json:"link_aggregation_group,omitempty"`
	HttpStatusCode       int                   `json:"-"`
}

func (o CreateLinkAggregationGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLinkAggregationGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateLinkAggregationGroupResponse", string(data)}, " ")
}
