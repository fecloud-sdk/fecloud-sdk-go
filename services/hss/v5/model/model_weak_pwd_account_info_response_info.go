package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type WeakPwdAccountInfoResponseInfo struct {
	UserName *string `json:"user_name,omitempty"`

	ServiceType *string `json:"service_type,omitempty"`

	Duration *int32 `json:"duration,omitempty"`
}

func (o WeakPwdAccountInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WeakPwdAccountInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"WeakPwdAccountInfoResponseInfo", string(data)}, " ")
}
