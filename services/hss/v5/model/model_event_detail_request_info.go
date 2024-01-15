package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type EventDetailRequestInfo struct {
	AgentId *string `json:"agent_id,omitempty"`

	ProcessPid *int32 `json:"process_pid,omitempty"`

	FileHash *string `json:"file_hash,omitempty"`

	FilePath *string `json:"file_path,omitempty"`

	FileAttr *string `json:"file_attr,omitempty"`

	Keyword *string `json:"keyword,omitempty"`

	Hash *string `json:"hash,omitempty"`

	PrivateIp *string `json:"private_ip,omitempty"`

	LoginIp *string `json:"login_ip,omitempty"`

	LoginUserName *string `json:"login_user_name,omitempty"`
}

func (o EventDetailRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventDetailRequestInfo struct{}"
	}

	return strings.Join([]string{"EventDetailRequestInfo", string(data)}, " ")
}
