package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateLoadBalancerRequest struct {
	LoadbalancerId string `json:"loadbalancer_id"`

	Body *UpdateLoadBalancerRequestBody `json:"body,omitempty"`
}

func (o UpdateLoadBalancerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLoadBalancerRequest struct{}"
	}

	return strings.Join([]string{"UpdateLoadBalancerRequest", string(data)}, " ")
}
