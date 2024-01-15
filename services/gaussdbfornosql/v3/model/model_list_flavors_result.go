package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListFlavorsResult struct {
	EngineName string `json:"engine_name"`

	EngineVersion string `json:"engine_version"`

	Vcpus string `json:"vcpus"`

	Ram string `json:"ram"`

	SpecCode string `json:"spec_code"`

	AvailabilityZone []string `json:"availability_zone"`

	AzStatus *interface{} `json:"az_status"`
}

func (o ListFlavorsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsResult struct{}"
	}

	return strings.Join([]string{"ListFlavorsResult", string(data)}, " ")
}
