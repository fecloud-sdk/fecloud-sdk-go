package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListDedicatedResourceResult struct {
	Id string `json:"id"`

	ResourceName string `json:"resource_name"`

	EngineName string `json:"engine_name"`

	AvailabilityZone string `json:"availability_zone"`

	Architecture string `json:"architecture"`

	Capacity *DedicatedResourceCapacity `json:"capacity"`

	Status string `json:"status"`
}

func (o ListDedicatedResourceResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDedicatedResourceResult struct{}"
	}

	return strings.Join([]string{"ListDedicatedResourceResult", string(data)}, " ")
}
