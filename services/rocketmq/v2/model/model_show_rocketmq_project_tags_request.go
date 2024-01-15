package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowRocketmqProjectTagsRequest struct {
}

func (o ShowRocketmqProjectTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRocketmqProjectTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowRocketmqProjectTagsRequest", string(data)}, " ")
}
