package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteVpcTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteVpcTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteVpcTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteVpcTagsResponse", string(data)}, " ")
}
