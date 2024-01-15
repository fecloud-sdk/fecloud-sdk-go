package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListQueryProjectResourceTagsResponse struct {
	Tags           *[]TagValuesList `json:"tags,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListQueryProjectResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueryProjectResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"ListQueryProjectResourceTagsResponse", string(data)}, " ")
}
