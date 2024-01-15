package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateEndpointRoutetableRequest struct {
	VpcEndpointId string `json:"vpc_endpoint_id"`

	ContentType *string `json:"Content-Type,omitempty"`

	Body *UpdateEndpointRoutetableRequestBody `json:"body,omitempty"`
}

func (o UpdateEndpointRoutetableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointRoutetableRequest struct{}"
	}

	return strings.Join([]string{"UpdateEndpointRoutetableRequest", string(data)}, " ")
}
