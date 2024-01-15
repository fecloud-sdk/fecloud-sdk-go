package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListUserStatisticsResponse struct {
	TotalNum *int32 `json:"total_num,omitempty"`

	DataList       *[]UserStatisticInfoResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ListUserStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListUserStatisticsResponse", string(data)}, " ")
}
