package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowRocketmqProjectTagsResponse struct {
	Tags           *[]TagMultyValueEntity `json:"tags,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowRocketmqProjectTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRocketmqProjectTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowRocketmqProjectTagsResponse", string(data)}, " ")
}
