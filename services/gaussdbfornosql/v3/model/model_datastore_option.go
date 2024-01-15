package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DatastoreOption struct {
	Type string `json:"type"`

	Version string `json:"version"`

	StorageEngine string `json:"storage_engine"`
}

func (o DatastoreOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatastoreOption struct{}"
	}

	return strings.Join([]string{"DatastoreOption", string(data)}, " ")
}
