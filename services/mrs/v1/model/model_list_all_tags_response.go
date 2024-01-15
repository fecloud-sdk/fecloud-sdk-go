package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListAllTagsResponse struct {
	Tags           *[]TagWithMultiValue `json:"tags,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListAllTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllTagsResponse struct{}"
	}

	return strings.Join([]string{"ListAllTagsResponse", string(data)}, " ")
}
