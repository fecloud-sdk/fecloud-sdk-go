package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowInstanceConfigurationResponse struct {
	DatastoreVersionName *string `json:"datastore_version_name,omitempty"`

	DatastoreName *string `json:"datastore_name,omitempty"`

	Created *string `json:"created,omitempty"`

	Updated *string `json:"updated,omitempty"`

	Id *string `json:"id,omitempty"`

	Mode *string `json:"mode,omitempty"`

	ConfigurationParameters *[]ConfigurationParameterResult `json:"configuration_parameters,omitempty"`
	HttpStatusCode          int                             `json:"-"`
}

func (o ShowInstanceConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceConfigurationResponse", string(data)}, " ")
}
