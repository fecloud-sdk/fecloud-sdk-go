package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateLinkAggregationGroupRequest struct {
	DcLagId string `json:"dc_lag_id"`

	Body *UpdateLinkAggregationGroupRequestBody `json:"body,omitempty"`
}

func (o UpdateLinkAggregationGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLinkAggregationGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateLinkAggregationGroupRequest", string(data)}, " ")
}
