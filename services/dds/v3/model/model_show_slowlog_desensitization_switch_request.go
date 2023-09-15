package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowSlowlogDesensitizationSwitchRequest Request Object
type ShowSlowlogDesensitizationSwitchRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID，可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建
	InstanceId string `json:"instance_id"`
}

func (o ShowSlowlogDesensitizationSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlowlogDesensitizationSwitchRequest struct{}"
	}

	return strings.Join([]string{"ShowSlowlogDesensitizationSwitchRequest", string(data)}, " ")
}
