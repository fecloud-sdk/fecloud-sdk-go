package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListDedicatedResourcesRequest struct {
	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDedicatedResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDedicatedResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListDedicatedResourcesRequest", string(data)}, " ")
}
