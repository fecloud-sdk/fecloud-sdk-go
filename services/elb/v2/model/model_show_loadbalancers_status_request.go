package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowLoadbalancersStatusRequest struct {
	LoadbalancerId string `json:"loadbalancer_id"`
}

func (o ShowLoadbalancersStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLoadbalancersStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowLoadbalancersStatusRequest", string(data)}, " ")
}
