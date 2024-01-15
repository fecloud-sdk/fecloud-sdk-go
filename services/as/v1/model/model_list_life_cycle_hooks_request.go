package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListLifeCycleHooksRequest struct {
	ScalingGroupId string `json:"scaling_group_id"`
}

func (o ListLifeCycleHooksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLifeCycleHooksRequest struct{}"
	}

	return strings.Join([]string{"ListLifeCycleHooksRequest", string(data)}, " ")
}
