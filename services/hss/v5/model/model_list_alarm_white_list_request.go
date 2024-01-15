package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListAlarmWhiteListRequest struct {
	Region string `json:"region"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Hash *string `json:"hash,omitempty"`

	EventType *int32 `json:"event_type,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAlarmWhiteListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmWhiteListRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmWhiteListRequest", string(data)}, " ")
}
