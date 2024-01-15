package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowShardingBalancerResponse struct {
	IsOpen *bool `json:"is_open,omitempty"`

	ActiveWindow   *BalancerActiveWindow `json:"active_window,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowShardingBalancerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowShardingBalancerResponse struct{}"
	}

	return strings.Join([]string{"ShowShardingBalancerResponse", string(data)}, " ")
}
