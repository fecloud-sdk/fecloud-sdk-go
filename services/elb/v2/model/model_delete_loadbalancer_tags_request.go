package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteLoadbalancerTagsRequest struct {
	LoadbalancerId string `json:"loadbalancer_id"`

	Key string `json:"key"`
}

func (o DeleteLoadbalancerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLoadbalancerTagsRequest struct{}"
	}

	return strings.Join([]string{"DeleteLoadbalancerTagsRequest", string(data)}, " ")
}
