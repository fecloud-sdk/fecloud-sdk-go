package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// AutoscanConfigRequest 设置自动分析配置返回体
type AutoscanConfigRequest struct {

	// 是否开启自动分析
	EnableAutoScan *bool `json:"enable_auto_scan,omitempty"`

	// 每日分析时间，时间格式为21:00，时间为UTC时间(目前只支持每天一次)
	ScheduleAt *[]string `json:"schedule_at,omitempty"`
}

func (o AutoscanConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoscanConfigRequest struct{}"
	}

	return strings.Join([]string{"AutoscanConfigRequest", string(data)}, " ")
}
