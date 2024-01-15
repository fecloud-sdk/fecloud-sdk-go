package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateEndpointServiceRequest struct {
	ContentType *string `json:"Content-Type,omitempty"`

	Body *CreateEndpointServiceRequestBody `json:"body,omitempty"`
}

func (o CreateEndpointServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEndpointServiceRequest struct{}"
	}

	return strings.Join([]string{"CreateEndpointServiceRequest", string(data)}, " ")
}
