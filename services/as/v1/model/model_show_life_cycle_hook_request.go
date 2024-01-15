package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowLifeCycleHookRequest struct {
	ScalingGroupId string `json:"scaling_group_id"`

	LifecycleHookName string `json:"lifecycle_hook_name"`
}

func (o ShowLifeCycleHookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLifeCycleHookRequest struct{}"
	}

	return strings.Join([]string{"ShowLifeCycleHookRequest", string(data)}, " ")
}
