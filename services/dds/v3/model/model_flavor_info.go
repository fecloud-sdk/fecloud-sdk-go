package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type FlavorInfo struct {
	EngineName string `json:"engine_name"`

	Type string `json:"type"`

	Vcpus string `json:"vcpus"`

	Ram string `json:"ram"`

	SpecCode string `json:"spec_code"`

	AzStatus *interface{} `json:"az_status"`

	EngineVersions []string `json:"engine_versions"`
}

func (o FlavorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorInfo struct{}"
	}

	return strings.Join([]string{"FlavorInfo", string(data)}, " ")
}
