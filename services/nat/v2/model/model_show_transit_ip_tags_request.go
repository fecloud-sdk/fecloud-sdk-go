package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowTransitIpTagsRequest Request Object
type ShowTransitIpTagsRequest struct {

	// 中转IP的ID。
	ResourceId string `json:"resource_id"`
}

func (o ShowTransitIpTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTransitIpTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowTransitIpTagsRequest", string(data)}, " ")
}
