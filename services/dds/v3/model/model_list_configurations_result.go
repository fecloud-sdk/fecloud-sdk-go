package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListConfigurationsResult struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Description string `json:"description"`

	DatastoreVersion string `json:"datastore_version"`

	DatastoreName string `json:"datastore_name"`

	NodeType string `json:"node_type"`

	Created string `json:"created"`

	Updated string `json:"updated"`

	UserDefined bool `json:"user_defined"`
}

func (o ListConfigurationsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationsResult struct{}"
	}

	return strings.Join([]string{"ListConfigurationsResult", string(data)}, " ")
}
