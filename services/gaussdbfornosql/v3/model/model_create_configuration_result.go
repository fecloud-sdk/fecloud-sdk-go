package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateConfigurationResult struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	DatastoreVersionName string `json:"datastore_version_name"`

	DatastoreName string `json:"datastore_name"`

	Created string `json:"created"`

	Updated string `json:"updated"`
}

func (o CreateConfigurationResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigurationResult struct{}"
	}

	return strings.Join([]string{"CreateConfigurationResult", string(data)}, " ")
}
