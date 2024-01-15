package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AvailabilityZoneDetail struct {
	PrimaryAvailabilityZone string `json:"primary_availability_zone"`

	SecondaryAvailabilityZone string `json:"secondary_availability_zone"`
}

func (o AvailabilityZoneDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailabilityZoneDetail struct{}"
	}

	return strings.Join([]string{"AvailabilityZoneDetail", string(data)}, " ")
}
