package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type StatusResp struct {
	Loadbalancer *LoadbalancerInStatusResp `json:"loadbalancer"`
}

func (o StatusResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatusResp struct{}"
	}

	return strings.Join([]string{"StatusResp", string(data)}, " ")
}
