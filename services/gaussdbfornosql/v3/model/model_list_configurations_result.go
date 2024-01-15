package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListConfigurationsResult struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	DatastoreVersionName string `json:"datastore_version_name"`

	DatastoreName string `json:"datastore_name"`

	Created string `json:"created"`

	Updated string `json:"updated"`

	Mode string `json:"mode"`

	UserDefined bool `json:"user_defined"`
}

func (o ListConfigurationsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationsResult struct{}"
	}

	return strings.Join([]string{"ListConfigurationsResult", string(data)}, " ")
}
