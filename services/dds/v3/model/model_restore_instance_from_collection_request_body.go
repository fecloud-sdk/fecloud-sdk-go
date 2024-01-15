package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RestoreInstanceFromCollectionRequestBody struct {
	RestoreCollections []RestoreInstanceFromCollectionRequestBodyRestoreCollections `json:"restore_collections"`
}

func (o RestoreInstanceFromCollectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreInstanceFromCollectionRequestBody struct{}"
	}

	return strings.Join([]string{"RestoreInstanceFromCollectionRequestBody", string(data)}, " ")
}
