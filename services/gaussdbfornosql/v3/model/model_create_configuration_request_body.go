package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateConfigurationRequestBody struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	Datastore *CreateConfigurationDatastoreOption `json:"datastore"`

	Values map[string]string `json:"values,omitempty"`
}

func (o CreateConfigurationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateConfigurationRequestBody", string(data)}, " ")
}
