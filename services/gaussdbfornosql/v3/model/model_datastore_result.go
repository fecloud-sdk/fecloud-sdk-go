package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DatastoreResult struct {
	Type *string `json:"type,omitempty"`

	Version *string `json:"version,omitempty"`

	StorageEngine *string `json:"storage_engine,omitempty"`
}

func (o DatastoreResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatastoreResult struct{}"
	}

	return strings.Join([]string{"DatastoreResult", string(data)}, " ")
}
