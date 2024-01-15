package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateEndpointServicePermissionDescRequest struct {
	ContentType *string `json:"Content-Type,omitempty"`

	VpcEndpointServiceId string `json:"vpc_endpoint_service_id"`

	PermissionId string `json:"permission_id"`

	Body *UpdatePermissionDescRequest `json:"body,omitempty"`
}

func (o UpdateEndpointServicePermissionDescRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointServicePermissionDescRequest struct{}"
	}

	return strings.Join([]string{"UpdateEndpointServicePermissionDescRequest", string(data)}, " ")
}
