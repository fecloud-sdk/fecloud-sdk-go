package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListImageTagsResponse struct {
	Tags           *[]ResourceTag `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListImageTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageTagsResponse struct{}"
	}

	return strings.Join([]string{"ListImageTagsResponse", string(data)}, " ")
}
