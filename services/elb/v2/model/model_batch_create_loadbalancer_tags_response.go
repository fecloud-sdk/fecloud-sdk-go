package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateLoadbalancerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateLoadbalancerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateLoadbalancerTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateLoadbalancerTagsResponse", string(data)}, " ")
}
