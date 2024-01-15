package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateEndpointPolicyRequest struct {
	VpcEndpointId string `json:"vpc_endpoint_id"`

	ContentType *string `json:"Content-Type,omitempty"`

	Body *UpdateEndpointPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateEndpointPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateEndpointPolicyRequest", string(data)}, " ")
}
