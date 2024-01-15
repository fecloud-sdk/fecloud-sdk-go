package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RouterWithStatus struct {
	Status *string `json:"status,omitempty"`

	RouterId *string `json:"router_id,omitempty"`

	RouterRegion *string `json:"router_region,omitempty"`
}

func (o RouterWithStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RouterWithStatus struct{}"
	}

	return strings.Join([]string{"RouterWithStatus", string(data)}, " ")
}
