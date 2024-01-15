package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListInstancesDatastoreResult struct {
	Type string `json:"type"`

	Version string `json:"version"`

	PatchAvailable bool `json:"patch_available"`

	WholeVersion string `json:"whole_version"`
}

func (o ListInstancesDatastoreResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesDatastoreResult struct{}"
	}

	return strings.Join([]string{"ListInstancesDatastoreResult", string(data)}, " ")
}
