package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowLoadBalancerRequest struct {
	LoadbalancerId string `json:"loadbalancer_id"`
}

func (o ShowLoadBalancerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLoadBalancerRequest struct{}"
	}

	return strings.Join([]string{"ShowLoadBalancerRequest", string(data)}, " ")
}
