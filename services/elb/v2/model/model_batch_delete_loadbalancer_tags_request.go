package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// BatchDeleteLoadbalancerTagsRequest Request Object
type BatchDeleteLoadbalancerTagsRequest struct {

	// 负载均衡器ID。
	LoadbalancerId string `json:"loadbalancer_id"`

	Body *BatchDeleteLoadbalancerTagsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteLoadbalancerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteLoadbalancerTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteLoadbalancerTagsRequest", string(data)}, " ")
}
