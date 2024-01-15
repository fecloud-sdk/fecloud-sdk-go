package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type WeakPwdListInfoResponseInfo struct {
	HostId *string `json:"host_id,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	HostIp *string `json:"host_ip,omitempty"`

	WeakPwdAccounts *[]WeakPwdAccountInfoResponseInfo `json:"weak_pwd_accounts,omitempty"`
}

func (o WeakPwdListInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WeakPwdListInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"WeakPwdListInfoResponseInfo", string(data)}, " ")
}
