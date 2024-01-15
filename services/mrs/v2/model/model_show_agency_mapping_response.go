package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowAgencyMappingResponse struct {
	AgencyMappings *[]AgencyMapping `json:"agency_mappings,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowAgencyMappingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgencyMappingResponse struct{}"
	}

	return strings.Join([]string{"ShowAgencyMappingResponse", string(data)}, " ")
}
