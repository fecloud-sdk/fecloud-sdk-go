package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListBackgroundTasksRequest struct {
	InstanceId string `json:"instance_id"`

	Start *int32 `json:"start,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	BeginTime *string `json:"begin_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`
}

func (o ListBackgroundTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackgroundTasksRequest struct{}"
	}

	return strings.Join([]string{"ListBackgroundTasksRequest", string(data)}, " ")
}
