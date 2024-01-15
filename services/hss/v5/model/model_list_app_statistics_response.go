package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListAppStatisticsResponse struct {
	TotalNum *int32 `json:"total_num,omitempty"`

	DataList       *[]AppStatisticResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListAppStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListAppStatisticsResponse", string(data)}, " ")
}
