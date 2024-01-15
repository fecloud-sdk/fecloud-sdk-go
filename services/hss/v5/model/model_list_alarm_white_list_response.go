package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListAlarmWhiteListResponse struct {
	TotalNum *int32 `json:"total_num,omitempty"`

	EventTypeList *[]int32 `json:"event_type_list,omitempty"`

	DataList       *[]AlarmWhiteListResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListAlarmWhiteListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmWhiteListResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmWhiteListResponse", string(data)}, " ")
}
