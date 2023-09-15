package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// UpdateLifeCycleHookRequest Request Object
type UpdateLifeCycleHookRequest struct {

	// 伸缩组标识。
	ScalingGroupId string `json:"scaling_group_id"`

	// 生命周期挂钩标识。
	LifecycleHookName string `json:"lifecycle_hook_name"`

	Body *UpdateLifeCycleHookOption `json:"body,omitempty"`
}

func (o UpdateLifeCycleHookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLifeCycleHookRequest struct{}"
	}

	return strings.Join([]string{"UpdateLifeCycleHookRequest", string(data)}, " ")
}
