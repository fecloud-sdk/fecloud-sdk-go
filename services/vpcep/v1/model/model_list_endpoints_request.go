package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListEndpointsRequest struct {
	ContentType *string `json:"Content-Type,omitempty"`

	EndpointServiceName *string `json:"endpoint_service_name,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	Id *string `json:"id,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	SortKey *string `json:"sort_key,omitempty"`

	SortDir *string `json:"sort_dir,omitempty"`

	PublicBorderGroup *string `json:"public_border_group,omitempty"`
}

func (o ListEndpointsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointsRequest struct{}"
	}

	return strings.Join([]string{"ListEndpointsRequest", string(data)}, " ")
}
