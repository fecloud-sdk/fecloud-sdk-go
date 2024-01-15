package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteLinkAggregationGroupRequest struct {
	DcLagId string `json:"dc_lag_id"`
}

func (o DeleteLinkAggregationGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLinkAggregationGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteLinkAggregationGroupRequest", string(data)}, " ")
}
