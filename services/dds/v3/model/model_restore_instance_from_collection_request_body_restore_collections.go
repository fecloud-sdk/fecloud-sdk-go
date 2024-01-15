package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RestoreInstanceFromCollectionRequestBodyRestoreCollections struct {
	Database string `json:"database"`

	RestoreDatabaseTime *string `json:"restore_database_time,omitempty"`

	Collections *[]RestoreInstanceFromCollectionRequestBodyCollections `json:"collections,omitempty"`
}

func (o RestoreInstanceFromCollectionRequestBodyRestoreCollections) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreInstanceFromCollectionRequestBodyRestoreCollections struct{}"
	}

	return strings.Join([]string{"RestoreInstanceFromCollectionRequestBodyRestoreCollections", string(data)}, " ")
}
