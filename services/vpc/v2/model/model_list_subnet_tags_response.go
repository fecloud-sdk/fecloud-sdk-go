package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListSubnetTagsResponse struct {
	Tags           *[]ListTag `json:"tags,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListSubnetTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubnetTagsResponse struct{}"
	}

	return strings.Join([]string{"ListSubnetTagsResponse", string(data)}, " ")
}
