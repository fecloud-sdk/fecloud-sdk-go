package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowLoadBalancerStatusRequest struct {
	LoadbalancerId string `json:"loadbalancer_id"`
}

func (o ShowLoadBalancerStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLoadBalancerStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowLoadBalancerStatusRequest", string(data)}, " ")
}
