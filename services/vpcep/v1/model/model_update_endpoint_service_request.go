package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateEndpointServiceRequest struct {
	ContentType *string `json:"Content-Type,omitempty"`

	VpcEndpointServiceId string `json:"vpc_endpoint_service_id"`

	Body *UpdateEndpointServiceRequestBody `json:"body,omitempty"`
}

func (o UpdateEndpointServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointServiceRequest struct{}"
	}

	return strings.Join([]string{"UpdateEndpointServiceRequest", string(data)}, " ")
}
