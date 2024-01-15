package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateEndpointWhiteRequest struct {
	VpcEndpointId string `json:"vpc_endpoint_id"`

	ContentType *string `json:"Content-Type,omitempty"`

	Body *UpdateEndpointWhiteRequestBody `json:"body,omitempty"`
}

func (o UpdateEndpointWhiteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointWhiteRequest struct{}"
	}

	return strings.Join([]string{"UpdateEndpointWhiteRequest", string(data)}, " ")
}
