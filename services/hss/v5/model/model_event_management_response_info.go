package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type EventManagementResponseInfo struct {
	EventId *string `json:"event_id,omitempty"`

	EventClassId *string `json:"event_class_id,omitempty"`

	EventType *int32 `json:"event_type,omitempty"`

	EventName *string `json:"event_name,omitempty"`

	Severity *string `json:"severity,omitempty"`

	ContainerName *string `json:"container_name,omitempty"`

	ImageName *string `json:"image_name,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	HostId *string `json:"host_id,omitempty"`

	PrivateIp *string `json:"private_ip,omitempty"`

	PublicIp *string `json:"public_ip,omitempty"`

	OsType *string `json:"os_type,omitempty"`

	HostStatus *string `json:"host_status,omitempty"`

	AgentStatus *string `json:"agent_status,omitempty"`

	ProtectStatus *string `json:"protect_status,omitempty"`

	AssetValue *string `json:"asset_value,omitempty"`

	AttackPhase *string `json:"attack_phase,omitempty"`

	AttackTag *string `json:"attack_tag,omitempty"`

	OccurTime *int64 `json:"occur_time,omitempty"`

	HandleTime *int64 `json:"handle_time,omitempty"`

	HandleStatus *string `json:"handle_status,omitempty"`

	HandleMethod *string `json:"handle_method,omitempty"`

	Handler *string `json:"handler,omitempty"`

	OperateAcceptList *[]string `json:"operate_accept_list,omitempty"`

	OperateDetailList *[]EventDetailResponseInfo `json:"operate_detail_list,omitempty"`

	ForensicInfo *interface{} `json:"forensic_info,omitempty"`

	ResourceInfo *EventResourceResponseInfo `json:"resource_info,omitempty"`

	GeoInfo *interface{} `json:"geo_info,omitempty"`

	MalwareInfo *interface{} `json:"malware_info,omitempty"`

	NetworkInfo *interface{} `json:"network_info,omitempty"`

	AppInfo *interface{} `json:"app_info,omitempty"`

	SystemInfo *interface{} `json:"system_info,omitempty"`

	ExtendInfo *interface{} `json:"extend_info,omitempty"`

	Recommendation *string `json:"recommendation,omitempty"`

	ProcessInfoList *[]EventProcessResponseInfo `json:"process_info_list,omitempty"`

	UserInfoList *[]EventUserResponseInfo `json:"user_info_list,omitempty"`

	FileInfoList *[]EventFileResponseInfo `json:"file_info_list,omitempty"`

	EventDetails *string `json:"event_details,omitempty"`
}

func (o EventManagementResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventManagementResponseInfo struct{}"
	}

	return strings.Join([]string{"EventManagementResponseInfo", string(data)}, " ")
}
