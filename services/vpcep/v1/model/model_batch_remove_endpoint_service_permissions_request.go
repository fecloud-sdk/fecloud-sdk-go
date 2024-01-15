package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchRemoveEndpointServicePermissionsRequest struct {
	ContentType *string `json:"Content-Type,omitempty"`

	VpcEndpointServiceId string `json:"vpc_endpoint_service_id"`

	Body *BatchRemovePermissionRequest `json:"body,omitempty"`
}

func (o BatchRemoveEndpointServicePermissionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRemoveEndpointServicePermissionsRequest struct{}"
	}

	return strings.Join([]string{"BatchRemoveEndpointServicePermissionsRequest", string(data)}, " ")
}
