package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateConfigurationRequestBody struct {
	Name string `json:"name"`

	Description string `json:"description"`

	ParameterValues map[string]string `json:"parameter_values"`

	Datastore *DatastoreResult `json:"datastore"`
}

func (o CreateConfigurationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateConfigurationRequestBody", string(data)}, " ")
}
