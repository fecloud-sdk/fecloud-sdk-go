package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteLoadbalancerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteLoadbalancerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteLoadbalancerTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteLoadbalancerTagsResponse", string(data)}, " ")
}
