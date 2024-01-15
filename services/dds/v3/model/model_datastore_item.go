package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DatastoreItem struct {
	Type string `json:"type"`

	Version string `json:"version"`

	PatchAvailable bool `json:"patch_available"`
}

func (o DatastoreItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatastoreItem struct{}"
	}

	return strings.Join([]string{"DatastoreItem", string(data)}, " ")
}
