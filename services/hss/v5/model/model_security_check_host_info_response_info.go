package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SecurityCheckHostInfoResponseInfo struct {
	HostId *string `json:"host_id,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	HostPublicIp *string `json:"host_public_ip,omitempty"`

	HostPrivateIp *string `json:"host_private_ip,omitempty"`

	ScanTime *int64 `json:"scan_time,omitempty"`

	FailedNum *int32 `json:"failed_num,omitempty"`

	PassedNum *int32 `json:"passed_num,omitempty"`
}

func (o SecurityCheckHostInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityCheckHostInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"SecurityCheckHostInfoResponseInfo", string(data)}, " ")
}
