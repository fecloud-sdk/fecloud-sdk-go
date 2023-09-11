package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateVpcTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateVpcTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateVpcTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateVpcTagsResponse", string(data)}, " ")
}
