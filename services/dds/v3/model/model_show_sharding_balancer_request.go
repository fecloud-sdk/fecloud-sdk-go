package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowShardingBalancerRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o ShowShardingBalancerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowShardingBalancerRequest struct{}"
	}

	return strings.Join([]string{"ShowShardingBalancerRequest", string(data)}, " ")
}
