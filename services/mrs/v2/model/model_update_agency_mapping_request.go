package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateAgencyMappingRequest struct {
	ClusterId string `json:"cluster_id"`

	Body *AgencyMappingArray `json:"body,omitempty"`
}

func (o UpdateAgencyMappingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAgencyMappingRequest struct{}"
	}

	return strings.Join([]string{"UpdateAgencyMappingRequest", string(data)}, " ")
}
