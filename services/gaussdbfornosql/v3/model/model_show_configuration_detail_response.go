package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowConfigurationDetailResponse struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	DatastoreVersionName *string `json:"datastore_version_name,omitempty"`

	DatastoreName *string `json:"datastore_name,omitempty"`

	Created *string `json:"created,omitempty"`

	Updated *string `json:"updated,omitempty"`

	ConfigurationParameters *[]ConfigurationParameterResult `json:"configuration_parameters,omitempty"`
	HttpStatusCode          int                             `json:"-"`
}

func (o ShowConfigurationDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigurationDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowConfigurationDetailResponse", string(data)}, " ")
}
