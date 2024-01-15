package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListAvailabilityZonesResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	AvailabilityZones *[][]AvailabilityZone `json:"availability_zones,omitempty"`
	HttpStatusCode    int                   `json:"-"`
}

func (o ListAvailabilityZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailabilityZonesResponse struct{}"
	}

	return strings.Join([]string{"ListAvailabilityZonesResponse", string(data)}, " ")
}
