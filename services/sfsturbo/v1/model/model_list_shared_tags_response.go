package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListSharedTagsResponse struct {
	Tags           *[]Tag `json:"tags,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSharedTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSharedTagsResponse struct{}"
	}

	return strings.Join([]string{"ListSharedTagsResponse", string(data)}, " ")
}
