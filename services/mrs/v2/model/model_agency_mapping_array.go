package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AgencyMappingArray struct {
	AgencyMappings []AgencyMapping `json:"agency_mappings"`
}

func (o AgencyMappingArray) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyMappingArray struct{}"
	}

	return strings.Join([]string{"AgencyMappingArray", string(data)}, " ")
}
