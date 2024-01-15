package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListDatastoreVersionsResponse struct {
	Versions       *[]string `json:"versions,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListDatastoreVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatastoreVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListDatastoreVersionsResponse", string(data)}, " ")
}
