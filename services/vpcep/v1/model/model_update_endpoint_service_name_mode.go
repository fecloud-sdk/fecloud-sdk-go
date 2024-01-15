package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateEndpointServiceNameMode struct {
	EndpointServiceName *string `json:"endpoint_service_name,omitempty"`
}

func (o UpdateEndpointServiceNameMode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointServiceNameMode struct{}"
	}

	return strings.Join([]string{"UpdateEndpointServiceNameMode", string(data)}, " ")
}
