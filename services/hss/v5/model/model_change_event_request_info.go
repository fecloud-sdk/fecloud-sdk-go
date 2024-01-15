package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ChangeEventRequestInfo struct {
	OperateType string `json:"operate_type"`

	Handler *string `json:"handler,omitempty"`

	OperateEventList []OperateEventRequestInfo `json:"operate_event_list"`
}

func (o ChangeEventRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeEventRequestInfo struct{}"
	}

	return strings.Join([]string{"ChangeEventRequestInfo", string(data)}, " ")
}
