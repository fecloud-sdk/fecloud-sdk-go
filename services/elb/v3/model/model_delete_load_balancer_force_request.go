package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteLoadBalancerForceRequest struct {
	LoadbalancerId string `json:"loadbalancer_id"`
}

func (o DeleteLoadBalancerForceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLoadBalancerForceRequest struct{}"
	}

	return strings.Join([]string{"DeleteLoadBalancerForceRequest", string(data)}, " ")
}
