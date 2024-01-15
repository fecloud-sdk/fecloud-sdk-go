package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteEndpointRequest struct {
	ContentType *string `json:"Content-Type,omitempty"`

	VpcEndpointId string `json:"vpc_endpoint_id"`
}

func (o DeleteEndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEndpointRequest struct{}"
	}

	return strings.Join([]string{"DeleteEndpointRequest", string(data)}, " ")
}
