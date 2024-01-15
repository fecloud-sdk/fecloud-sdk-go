package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateLoadbalancerRequest struct {
	LoadbalancerId string `json:"loadbalancer_id"`

	Body *UpdateLoadbalancerRequestBody `json:"body,omitempty"`
}

func (o UpdateLoadbalancerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLoadbalancerRequest struct{}"
	}

	return strings.Join([]string{"UpdateLoadbalancerRequest", string(data)}, " ")
}
