package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AvailabilityZone struct {
	Code string `json:"code"`

	State string `json:"state"`

	Protocol []string `json:"protocol"`

	PublicBorderGroup string `json:"public_border_group"`

	Category int32 `json:"category"`
}

func (o AvailabilityZone) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailabilityZone struct{}"
	}

	return strings.Join([]string{"AvailabilityZone", string(data)}, " ")
}
