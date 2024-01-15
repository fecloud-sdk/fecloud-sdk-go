package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CloseProtectionInfoRequestInfo struct {
	HostIdList []string `json:"host_id_list"`

	AgentIdList []string `json:"agent_id_list"`

	CloseProtectionType string `json:"close_protection_type"`
}

func (o CloseProtectionInfoRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloseProtectionInfoRequestInfo struct{}"
	}

	return strings.Join([]string{"CloseProtectionInfoRequestInfo", string(data)}, " ")
}
