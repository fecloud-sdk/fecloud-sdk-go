package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListPortStatisticsResponse struct {
	TotalNum *int32 `json:"total_num,omitempty"`

	DataList       *[]PortStatisticResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListPortStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPortStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListPortStatisticsResponse", string(data)}, " ")
}
