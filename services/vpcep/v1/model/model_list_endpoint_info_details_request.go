package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListEndpointInfoDetailsRequest struct {
	ContentType *string `json:"Content-Type,omitempty"`

	VpcEndpointId string `json:"vpc_endpoint_id"`
}

func (o ListEndpointInfoDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointInfoDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListEndpointInfoDetailsRequest", string(data)}, " ")
}
