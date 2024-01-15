package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowEntityConfigurationResponse struct {
	DatastoreVersion *string `json:"datastore_version,omitempty"`

	DatastoreName *string `json:"datastore_name,omitempty"`

	Created *string `json:"created,omitempty"`

	Updated *string `json:"updated,omitempty"`

	Parameters     *[]EntityConfigurationParametersResult `json:"parameters,omitempty"`
	HttpStatusCode int                                    `json:"-"`
}

func (o ShowEntityConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEntityConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ShowEntityConfigurationResponse", string(data)}, " ")
}
