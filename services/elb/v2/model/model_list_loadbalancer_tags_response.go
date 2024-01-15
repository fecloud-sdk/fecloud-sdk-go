package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListLoadbalancerTagsResponse struct {
	Tags           *[]ListTag `json:"tags,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListLoadbalancerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLoadbalancerTagsResponse struct{}"
	}

	return strings.Join([]string{"ListLoadbalancerTagsResponse", string(data)}, " ")
}
