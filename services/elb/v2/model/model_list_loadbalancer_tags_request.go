package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListLoadbalancerTagsRequest struct {
}

func (o ListLoadbalancerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLoadbalancerTagsRequest struct{}"
	}

	return strings.Join([]string{"ListLoadbalancerTagsRequest", string(data)}, " ")
}
