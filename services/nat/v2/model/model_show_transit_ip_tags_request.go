package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowTransitIpTagsRequest struct {
	ResourceId string `json:"resource_id"`
}

func (o ShowTransitIpTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTransitIpTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowTransitIpTagsRequest", string(data)}, " ")
}
