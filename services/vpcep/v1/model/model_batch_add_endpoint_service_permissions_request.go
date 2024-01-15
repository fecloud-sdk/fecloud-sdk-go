package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchAddEndpointServicePermissionsRequest struct {
	ContentType *string `json:"Content-Type,omitempty"`

	VpcEndpointServiceId string `json:"vpc_endpoint_service_id"`

	Body *BatchAddPermissionRequest `json:"body,omitempty"`
}

func (o BatchAddEndpointServicePermissionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddEndpointServicePermissionsRequest struct{}"
	}

	return strings.Join([]string{"BatchAddEndpointServicePermissionsRequest", string(data)}, " ")
}
