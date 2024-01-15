package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListProcessStatisticsResponse struct {
	TotalNum *int32 `json:"total_num,omitempty"`

	DataList       *[]ProcessStatisticResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListProcessStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProcessStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListProcessStatisticsResponse", string(data)}, " ")
}
