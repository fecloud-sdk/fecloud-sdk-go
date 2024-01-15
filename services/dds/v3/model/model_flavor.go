package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Flavor struct {
	EngineName string `json:"engine_name"`

	Type string `json:"type"`

	Vcpus string `json:"vcpus"`

	Ram string `json:"ram"`

	SpecCode string `json:"spec_code"`

	AzStatus *interface{} `json:"az_status"`
}

func (o Flavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Flavor struct{}"
	}

	return strings.Join([]string{"Flavor", string(data)}, " ")
}
