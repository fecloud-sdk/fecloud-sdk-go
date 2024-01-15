package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RestoreInstanceFromCollectionRequestBodyCollections struct {
	OldName string `json:"old_name"`

	NewName *string `json:"new_name,omitempty"`

	RestoreCollectionTime string `json:"restore_collection_time"`
}

func (o RestoreInstanceFromCollectionRequestBodyCollections) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreInstanceFromCollectionRequestBodyCollections struct{}"
	}

	return strings.Join([]string{"RestoreInstanceFromCollectionRequestBodyCollections", string(data)}, " ")
}
