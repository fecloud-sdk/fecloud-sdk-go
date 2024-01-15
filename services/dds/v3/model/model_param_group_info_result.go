package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ParamGroupInfoResult struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Description string `json:"description"`

	DatastoreVersion string `json:"datastore_version"`

	DatastoreName string `json:"datastore_name"`

	Created string `json:"created"`

	Updated string `json:"updated"`
}

func (o ParamGroupInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParamGroupInfoResult struct{}"
	}

	return strings.Join([]string{"ParamGroupInfoResult", string(data)}, " ")
}
