package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteSubnetTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteSubnetTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSubnetTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteSubnetTagsResponse", string(data)}, " ")
}
