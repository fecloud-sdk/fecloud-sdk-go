package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowConfigurationParameterResponse struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	DatastoreVersion *string `json:"datastore_version,omitempty"`

	DatastoreName *string `json:"datastore_name,omitempty"`

	Description *string `json:"description,omitempty"`

	Created *string `json:"created,omitempty"`

	Updated *string `json:"updated,omitempty"`

	Parameters     *[]ConfigurationParametersResult `json:"parameters,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ShowConfigurationParameterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigurationParameterResponse struct{}"
	}

	return strings.Join([]string{"ShowConfigurationParameterResponse", string(data)}, " ")
}
