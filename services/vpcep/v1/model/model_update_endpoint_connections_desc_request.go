package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateEndpointConnectionsDescRequest struct {
	VpcEndpointServiceId string `json:"vpc_endpoint_service_id"`

	Body *UpdateEpConnections `json:"body,omitempty"`
}

func (o UpdateEndpointConnectionsDescRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointConnectionsDescRequest struct{}"
	}

	return strings.Join([]string{"UpdateEndpointConnectionsDescRequest", string(data)}, " ")
}
