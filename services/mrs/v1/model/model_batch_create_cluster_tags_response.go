package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateClusterTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateClusterTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateClusterTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateClusterTagsResponse", string(data)}, " ")
}
