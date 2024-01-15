package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteClusterTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteClusterTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteClusterTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteClusterTagsResponse", string(data)}, " ")
}
