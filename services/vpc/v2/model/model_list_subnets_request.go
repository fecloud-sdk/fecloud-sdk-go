package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListSubnetsRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`
}

func (o ListSubnetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubnetsRequest struct{}"
	}

	return strings.Join([]string{"ListSubnetsRequest", string(data)}, " ")
}
