package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListKmsTagsResponse struct {
	Tags           *[]Tag `json:"tags,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListKmsTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKmsTagsResponse struct{}"
	}

	return strings.Join([]string{"ListKmsTagsResponse", string(data)}, " ")
}
