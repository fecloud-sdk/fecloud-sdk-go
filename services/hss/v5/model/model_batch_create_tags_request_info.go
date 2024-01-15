package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateTagsRequestInfo struct {
	Tags *[]ResourceTagInfo `json:"tags,omitempty"`

	SysTags *[]ResourceTagInfo `json:"sys_tags,omitempty"`
}

func (o BatchCreateTagsRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateTagsRequestInfo struct{}"
	}

	return strings.Join([]string{"BatchCreateTagsRequestInfo", string(data)}, " ")
}
