package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateDeleteTransitIpTagsRequest struct {
	ResourceId string `json:"resource_id"`

	Body *BatchOperateResourceTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateDeleteTransitIpTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateDeleteTransitIpTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateDeleteTransitIpTagsRequest", string(data)}, " ")
}
