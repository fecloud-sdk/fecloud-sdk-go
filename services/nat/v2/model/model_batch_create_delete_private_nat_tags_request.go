package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateDeletePrivateNatTagsRequest struct {
	ResourceId string `json:"resource_id"`

	Body *BatchOperateResourceTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateDeletePrivateNatTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateDeletePrivateNatTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateDeletePrivateNatTagsRequest", string(data)}, " ")
}
