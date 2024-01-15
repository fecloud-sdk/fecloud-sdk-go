package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type OperateEventRequestInfo struct {
	EventClassId string `json:"event_class_id"`

	EventId string `json:"event_id"`

	EventType int32 `json:"event_type"`

	OccurTime int64 `json:"occur_time"`

	OperateDetailList []EventDetailRequestInfo `json:"operate_detail_list"`
}

func (o OperateEventRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateEventRequestInfo struct{}"
	}

	return strings.Join([]string{"OperateEventRequestInfo", string(data)}, " ")
}
