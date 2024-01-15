package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListAutoLaunchStatisticsResponse struct {
	TotalNum *int32 `json:"total_num,omitempty"`

	DataList       *[]AutoLaunchStatisticsResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ListAutoLaunchStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutoLaunchStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListAutoLaunchStatisticsResponse", string(data)}, " ")
}
