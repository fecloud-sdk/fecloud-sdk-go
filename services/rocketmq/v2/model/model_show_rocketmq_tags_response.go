package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowRocketmqTagsResponse struct {
	Tags           *[]TagEntity `json:"tags,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowRocketmqTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRocketmqTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowRocketmqTagsResponse", string(data)}, " ")
}
