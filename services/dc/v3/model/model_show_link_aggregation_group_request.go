package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowLinkAggregationGroupRequest struct {
	Fields *[]string `json:"fields,omitempty"`

	DcLagId string `json:"dc_lag_id"`
}

func (o ShowLinkAggregationGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLinkAggregationGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowLinkAggregationGroupRequest", string(data)}, " ")
}
