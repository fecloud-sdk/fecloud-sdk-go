package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteLoadbalancerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLoadbalancerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLoadbalancerTagsResponse struct{}"
	}

	return strings.Join([]string{"DeleteLoadbalancerTagsResponse", string(data)}, " ")
}
