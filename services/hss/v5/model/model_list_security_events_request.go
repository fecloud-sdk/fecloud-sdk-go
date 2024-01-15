package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListSecurityEventsRequest struct {
	Region string `json:"region"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	LastDays *int32 `json:"last_days,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	HostId *string `json:"host_id,omitempty"`

	PrivateIp *string `json:"private_ip,omitempty"`

	ContainerName *string `json:"container_name,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	EventTypes *[]int32 `json:"event_types,omitempty"`

	HandleStatus *string `json:"handle_status,omitempty"`

	Severity *string `json:"severity,omitempty"`

	Category string `json:"category"`

	BeginTime *string `json:"begin_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	EventClassIds *[]string `json:"event_class_ids,omitempty"`
}

func (o ListSecurityEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityEventsRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityEventsRequest", string(data)}, " ")
}
