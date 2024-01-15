package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteEndpointServiceRequest struct {
	ContentType *string `json:"Content-Type,omitempty"`

	VpcEndpointServiceId string `json:"vpc_endpoint_service_id"`
}

func (o DeleteEndpointServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEndpointServiceRequest struct{}"
	}

	return strings.Join([]string{"DeleteEndpointServiceRequest", string(data)}, " ")
}
