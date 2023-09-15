package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateLoadbalancerResponse Response Object
type CreateLoadbalancerResponse struct {
	Loadbalancer *LoadbalancerResp `json:"loadbalancer,omitempty"`

	// 订单号（包周期场景返回该字段）
	OrderId *string `json:"order_id,omitempty"`

	// 负载均衡器的id（包周期场景返回该字段）
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
