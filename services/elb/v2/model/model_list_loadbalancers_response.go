package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListLoadbalancersResponse struct {
	Loadbalancers  *[]LoadbalancerResp `json:"loadbalancers,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListLoadbalancersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLoadbalancersResponse struct{}"
	}

	return strings.Join([]string{"ListLoadbalancersResponse", string(data)}, " ")
}
