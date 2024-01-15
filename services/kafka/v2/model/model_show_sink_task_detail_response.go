package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowSinkTaskDetailResponse struct {
	TaskName *string `json:"task_name,omitempty"`

	DestinationType *string `json:"destination_type,omitempty"`

	CreateTime *int64 `json:"create_time,omitempty"`

	Status *string `json:"status,omitempty"`

	Topics *string `json:"topics,omitempty"`

	ObsDestinationDescriptor *ShowSinkTaskDetailRespObsDestinationDescriptor `json:"obs_destination_descriptor,omitempty"`

	TopicsInfo     *[]ShowSinkTaskDetailRespTopicsInfo `json:"topics_info,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ShowSinkTaskDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSinkTaskDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowSinkTaskDetailResponse", string(data)}, " ")
}
