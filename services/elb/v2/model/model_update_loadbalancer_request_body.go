package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateLoadbalancerRequestBody struct {
	Loadbalancer *UpdateLoadbalancerReq `json:"loadbalancer"`
}

func (o UpdateLoadbalancerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLoadbalancerRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateLoadbalancerRequestBody", string(data)}, " ")
}
