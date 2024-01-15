package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowLoadbalancerTagsResponse struct {
	Tags           *[]ResourceTag `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowLoadbalancerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLoadbalancerTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowLoadbalancerTagsResponse", string(data)}, " ")
}
