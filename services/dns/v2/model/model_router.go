package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Router struct {
	RouterId string `json:"router_id"`

	RouterRegion *string `json:"router_region,omitempty"`

	Status *string `json:"status,omitempty"`
}

func (o Router) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Router struct{}"
	}

	return strings.Join([]string{"Router", string(data)}, " ")
}
