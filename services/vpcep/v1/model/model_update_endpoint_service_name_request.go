package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateEndpointServiceNameRequest struct {
	VpcEndpointServiceId string `json:"vpc_endpoint_service_id"`

	Body *UpdateEndpointServiceNameMode `json:"body,omitempty"`
}

func (o UpdateEndpointServiceNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointServiceNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateEndpointServiceNameRequest", string(data)}, " ")
}
