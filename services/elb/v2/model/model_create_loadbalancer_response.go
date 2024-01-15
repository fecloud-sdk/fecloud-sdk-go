package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateLoadbalancerResponse struct {
	Loadbalancer *LoadbalancerResp `json:"loadbalancer,omitempty"`

	OrderId *string `json:"order_id,omitempty"`

	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateLoadbalancerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadbalancerResponse struct{}"
	}

	return strings.Join([]string{"CreateLoadbalancerResponse", string(data)}, " ")
}
