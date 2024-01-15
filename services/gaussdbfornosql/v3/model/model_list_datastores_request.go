package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListDatastoresRequest struct {
	DatastoreName string `json:"datastore_name"`
}

func (o ListDatastoresRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatastoresRequest struct{}"
	}

	return strings.Join([]string{"ListDatastoresRequest", string(data)}, " ")
}
