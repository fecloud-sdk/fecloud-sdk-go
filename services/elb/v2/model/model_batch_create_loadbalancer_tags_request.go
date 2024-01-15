package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateLoadbalancerTagsRequest struct {
	LoadbalancerId string `json:"loadbalancer_id"`

	Body *BatchCreateLoadbalancerTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateLoadbalancerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateLoadbalancerTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateLoadbalancerTagsRequest", string(data)}, " ")
}
