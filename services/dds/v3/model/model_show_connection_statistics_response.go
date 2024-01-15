package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowConnectionStatisticsResponse struct {
	TotalConnections *int32 `json:"total_connections,omitempty"`

	TotalInnerConnections *int32 `json:"total_inner_connections,omitempty"`

	TotalOuterConnections *int32 `json:"total_outer_connections,omitempty"`

	InnerConnections *[]QueryConnectionsResponse `json:"inner_connections,omitempty"`

	OuterConnections *[]QueryConnectionsResponse `json:"outer_connections,omitempty"`
	HttpStatusCode   int                         `json:"-"`
}

func (o ShowConnectionStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConnectionStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowConnectionStatisticsResponse", string(data)}, " ")
}
