package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteLoadbalancerTagsRequest Request Object
type DeleteLoadbalancerTagsRequest struct {

	// 负载均衡器ID。
	LoadbalancerId string `json:"loadbalancer_id"`

	// 待删除标签的key值
	Key string `json:"key"`
}

func (o DeleteLoadbalancerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLoadbalancerTagsRequest struct{}"
	}

	return strings.Join([]string{"DeleteLoadbalancerTagsRequest", string(data)}, " ")
}