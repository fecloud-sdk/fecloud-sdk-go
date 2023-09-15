package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// UpdateLoadbalancerRequestBody This is a auto create Body Object
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
