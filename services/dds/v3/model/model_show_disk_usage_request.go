package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowDiskUsageRequest Request Object
type ShowDiskUsageRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID，可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。
	InstanceId string `json:"instance_id"`
}

func (o ShowDiskUsageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDiskUsageRequest struct{}"
	}

	return strings.Join([]string{"ShowDiskUsageRequest", string(data)}, " ")
}
