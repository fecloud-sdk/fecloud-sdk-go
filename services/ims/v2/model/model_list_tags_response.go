package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListTagsResponse struct {
	Tags           *[]string `json:"tags,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsResponse struct{}"
	}

	return strings.Join([]string{"ListTagsResponse", string(data)}, " ")
}
