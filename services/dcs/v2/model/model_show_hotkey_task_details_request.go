package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowHotkeyTaskDetailsRequest Request Object
type ShowHotkeyTaskDetailsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 热key分析任务ID。
	HotkeyId string `json:"hotkey_id"`
}

func (o ShowHotkeyTaskDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHotkeyTaskDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowHotkeyTaskDetailsRequest", string(data)}, " ")
}