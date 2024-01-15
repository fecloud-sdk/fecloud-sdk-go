package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteEndpointPolicyRequest struct {
	VpcEndpointId string `json:"vpc_endpoint_id"`

	ContentType *string `json:"Content-Type,omitempty"`
}

func (o DeleteEndpointPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEndpointPolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteEndpointPolicyRequest", string(data)}, " ")
}
