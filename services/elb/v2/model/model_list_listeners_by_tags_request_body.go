package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListListenersByTagsRequestBody struct {
	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Action string `json:"action"`

	Matches *[]ActionMatch `json:"matches,omitempty"`

	Tags *[]ActionTag `json:"tags,omitempty"`
}

func (o ListListenersByTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListListenersByTagsRequestBody struct{}"
	}

	return strings.Join([]string{"ListListenersByTagsRequestBody", string(data)}, " ")
}
