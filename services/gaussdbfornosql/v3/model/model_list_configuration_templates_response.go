package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListConfigurationTemplatesResponse struct {
	Count *int32 `json:"count,omitempty"`

	Quota *int32 `json:"quota,omitempty"`

	Configurations *[]ListConfigurationsResult `json:"configurations,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListConfigurationTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListConfigurationTemplatesResponse", string(data)}, " ")
}
