package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateVpcTagsRequest struct {
	VpcId string `json:"vpc_id"`

	Body *BatchCreateVpcTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateVpcTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateVpcTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateVpcTagsRequest", string(data)}, " ")
}
