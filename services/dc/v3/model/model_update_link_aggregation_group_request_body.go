package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateLinkAggregationGroupRequestBody struct {
	LinkAggregationGroup *UpdateLinkAggregationGroup `json:"link_aggregation_group,omitempty"`
}

func (o UpdateLinkAggregationGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLinkAggregationGroupRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateLinkAggregationGroupRequestBody", string(data)}, " ")
}
