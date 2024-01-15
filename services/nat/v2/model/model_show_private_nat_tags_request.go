package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowPrivateNatTagsRequest struct {
	ResourceId string `json:"resource_id"`
}

func (o ShowPrivateNatTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivateNatTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowPrivateNatTagsRequest", string(data)}, " ")
}
