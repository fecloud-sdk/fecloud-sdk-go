package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// UpdateLoadbalancerResponse Response Object
type UpdateLoadbalancerResponse struct {
	Loadbalancer   *LoadbalancerResp `json:"loadbalancer,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o UpdateLoadbalancerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLoadbalancerResponse struct{}"
	}

	return strings.Join([]string{"UpdateLoadbalancerResponse", string(data)}, " ")
}
