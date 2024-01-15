package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListClusterTagsResponse struct {
	Tags           *[]TagPlain `json:"tags,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListClusterTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterTagsResponse struct{}"
	}

	return strings.Join([]string{"ListClusterTagsResponse", string(data)}, " ")
}
