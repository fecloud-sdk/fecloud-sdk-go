package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateEndpointRequest struct {
	ContentType *string `json:"Content-Type,omitempty"`

	Body *CreateEndpointRequestBody `json:"body,omitempty"`
}

func (o CreateEndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEndpointRequest struct{}"
	}

	return strings.Join([]string{"CreateEndpointRequest", string(data)}, " ")
}
