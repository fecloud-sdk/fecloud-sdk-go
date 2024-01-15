package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteLifecycleHookRequest struct {
	ScalingGroupId string `json:"scaling_group_id"`

	LifecycleHookName string `json:"lifecycle_hook_name"`
}

func (o DeleteLifecycleHookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLifecycleHookRequest struct{}"
	}

	return strings.Join([]string{"DeleteLifecycleHookRequest", string(data)}, " ")
}
