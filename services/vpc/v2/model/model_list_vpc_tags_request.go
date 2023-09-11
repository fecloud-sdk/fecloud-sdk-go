package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListVpcTagsRequest struct {
}

func (o ListVpcTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcTagsRequest struct{}"
	}

	return strings.Join([]string{"ListVpcTagsRequest", string(data)}, " ")
}
