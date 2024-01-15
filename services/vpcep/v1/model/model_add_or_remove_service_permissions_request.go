package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AddOrRemoveServicePermissionsRequest struct {
	ContentType *string `json:"Content-Type,omitempty"`

	VpcEndpointServiceId string `json:"vpc_endpoint_service_id"`

	Body *AddOrRemoveServicePermissionsRequestBody `json:"body,omitempty"`
}

func (o AddOrRemoveServicePermissionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddOrRemoveServicePermissionsRequest struct{}"
	}

	return strings.Join([]string{"AddOrRemoveServicePermissionsRequest", string(data)}, " ")
}
