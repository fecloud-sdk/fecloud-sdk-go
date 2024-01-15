package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListFlavorInfosRequest struct {
	EngineName *string `json:"engine_name,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListFlavorInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorInfosRequest struct{}"
	}

	return strings.Join([]string{"ListFlavorInfosRequest", string(data)}, " ")
}
