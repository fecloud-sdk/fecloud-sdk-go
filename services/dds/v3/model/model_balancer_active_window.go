package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BalancerActiveWindow struct {
	StartTime *string `json:"start_time,omitempty"`

	StopTime *string `json:"stop_time,omitempty"`
}

func (o BalancerActiveWindow) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BalancerActiveWindow struct{}"
	}

	return strings.Join([]string{"BalancerActiveWindow", string(data)}, " ")
}
