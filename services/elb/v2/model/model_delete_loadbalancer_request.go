package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteLoadbalancerRequest struct {
	LoadbalancerId string `json:"loadbalancer_id"`
}

func (o DeleteLoadbalancerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLoadbalancerRequest struct{}"
	}

	return strings.Join([]string{"DeleteLoadbalancerRequest", string(data)}, " ")
}
