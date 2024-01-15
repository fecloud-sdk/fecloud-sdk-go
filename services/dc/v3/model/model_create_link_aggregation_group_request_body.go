package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateLinkAggregationGroupRequestBody struct {
	LinkAggregationGroup *CreateLinkAggregationGroup `json:"link_aggregation_group,omitempty"`
}

func (o CreateLinkAggregationGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLinkAggregationGroupRequestBody struct{}"
	}

	return strings.Join([]string{"CreateLinkAggregationGroupRequestBody", string(data)}, " ")
}
