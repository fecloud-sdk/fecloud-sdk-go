package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AcceptOrRejectEndpointRequest struct {
	ContentType *string `json:"Content-Type,omitempty"`

	VpcEndpointServiceId string `json:"vpc_endpoint_service_id"`

	Body *AcceptOrRejectEndpointRequestBody `json:"body,omitempty"`
}

func (o AcceptOrRejectEndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptOrRejectEndpointRequest struct{}"
	}

	return strings.Join([]string{"AcceptOrRejectEndpointRequest", string(data)}, " ")
}
