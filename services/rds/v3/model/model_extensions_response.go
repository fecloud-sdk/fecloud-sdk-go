package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ExtensionsResponse struct {
	Name *string `json:"name,omitempty"`

	DatabaseName *string `json:"database_name,omitempty"`

	Version *string `json:"version,omitempty"`

	SharedPreloadLibraries *string `json:"shared_preload_libraries,omitempty"`

	Created *bool `json:"created,omitempty"`

	Description *string `json:"description,omitempty"`
}

func (o ExtensionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtensionsResponse struct{}"
	}

	return strings.Join([]string{"ExtensionsResponse", string(data)}, " ")
}
