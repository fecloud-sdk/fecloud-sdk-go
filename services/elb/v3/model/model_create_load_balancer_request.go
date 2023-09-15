package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateLoadBalancerRequest Request Object
type CreateLoadBalancerRequest struct {
	Body *CreateLoadBalancerRequestBody `json:"body,omitempty"`
}

func (o CreateLoadBalancerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadBalancerRequest struct{}"
	}

	return strings.Join([]string{"CreateLoadBalancerRequest", string(data)}, " ")
}
