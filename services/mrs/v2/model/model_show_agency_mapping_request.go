package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowAgencyMappingRequest struct {
	ClusterId string `json:"cluster_id"`
}

func (o ShowAgencyMappingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgencyMappingRequest struct{}"
	}

	return strings.Join([]string{"ShowAgencyMappingRequest", string(data)}, " ")
}
