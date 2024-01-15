package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateLoadbalancerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateLoadbalancerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadbalancerTagsResponse struct{}"
	}

	return strings.Join([]string{"CreateLoadbalancerTagsResponse", string(data)}, " ")
}
