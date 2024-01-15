package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type EventDetailResponseInfo struct {
	AgentId *string `json:"agent_id,omitempty"`

	ProcessPid *int32 `json:"process_pid,omitempty"`

	IsParent *bool `json:"is_parent,omitempty"`

	FileHash *string `json:"file_hash,omitempty"`

	FilePath *string `json:"file_path,omitempty"`

	FileAttr *string `json:"file_attr,omitempty"`

	PrivateIp *string `json:"private_ip,omitempty"`

	LoginIp *string `json:"login_ip,omitempty"`

	LoginUserName *string `json:"login_user_name,omitempty"`

	Keyword *string `json:"keyword,omitempty"`

	Hash *string `json:"hash,omitempty"`
}

func (o EventDetailResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventDetailResponseInfo struct{}"
	}

	return strings.Join([]string{"EventDetailResponseInfo", string(data)}, " ")
}
