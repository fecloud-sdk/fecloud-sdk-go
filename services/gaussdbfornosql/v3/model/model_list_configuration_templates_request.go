package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListConfigurationTemplatesRequest struct {
	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListConfigurationTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListConfigurationTemplatesRequest", string(data)}, " ")
}
