package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateConfigurationDatastoreOption struct {
	Type string `json:"type"`

	Version string `json:"version"`

	Mode *string `json:"mode,omitempty"`
}

func (o CreateConfigurationDatastoreOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigurationDatastoreOption struct{}"
	}

	return strings.Join([]string{"CreateConfigurationDatastoreOption", string(data)}, " ")
}
