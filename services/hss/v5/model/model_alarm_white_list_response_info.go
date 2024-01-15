package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AlarmWhiteListResponseInfo struct {
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	Hash *string `json:"hash,omitempty"`

	Description *string `json:"description,omitempty"`

	EventType *int32 `json:"event_type,omitempty"`

	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o AlarmWhiteListResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmWhiteListResponseInfo struct{}"
	}

	return strings.Join([]string{"AlarmWhiteListResponseInfo", string(data)}, " ")
}
