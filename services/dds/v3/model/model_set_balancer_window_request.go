package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SetBalancerWindowRequest struct {
	InstanceId string `json:"instance_id"`

	Body *BalancerActiveWindow `json:"body,omitempty"`
}

func (o SetBalancerWindowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetBalancerWindowRequest struct{}"
	}

	return strings.Join([]string{"SetBalancerWindowRequest", string(data)}, " ")
}
