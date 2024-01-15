package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteLinkAggregationGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLinkAggregationGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLinkAggregationGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteLinkAggregationGroupResponse", string(data)}, " ")
}
