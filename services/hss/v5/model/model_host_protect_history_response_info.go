package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type HostProtectHistoryResponseInfo struct {
	OccrTime *int64 `json:"occr_time,omitempty"`

	FilePath *string `json:"file_path,omitempty"`

	FileOperation *string `json:"file_operation,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	HostIp *string `json:"host_ip,omitempty"`

	ProcessId *string `json:"process_id,omitempty"`

	ProcessName *string `json:"process_name,omitempty"`

	ProcessCmd *string `json:"process_cmd,omitempty"`
}

func (o HostProtectHistoryResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostProtectHistoryResponseInfo struct{}"
	}

	return strings.Join([]string{"HostProtectHistoryResponseInfo", string(data)}, " ")
}
