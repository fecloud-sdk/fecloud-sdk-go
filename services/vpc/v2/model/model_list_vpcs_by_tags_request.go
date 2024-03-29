package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListVpcsByTagsRequest struct {
	Body *ListVpcsByTagsRequestBody `json:"body,omitempty"`
}

func (o ListVpcsByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcsByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListVpcsByTagsRequest", string(data)}, " ")
}
