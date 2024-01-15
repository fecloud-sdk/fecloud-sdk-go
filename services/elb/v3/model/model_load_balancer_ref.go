package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type LoadBalancerRef struct {
	Id *string `json:"id,omitempty"`
}

func (o LoadBalancerRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadBalancerRef struct{}"
	}

	return strings.Join([]string{"LoadBalancerRef", string(data)}, " ")
}
