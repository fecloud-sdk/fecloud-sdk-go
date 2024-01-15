package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type QueryResourceInstanceTagsBody struct {
	Tags *[]TagValuesList `json:"tags,omitempty"`

	TagsAny *[]TagValuesList `json:"tags_any,omitempty"`

	NotTags *[]TagValuesList `json:"not_tags,omitempty"`

	NotTagsAny *[]TagValuesList `json:"not_tags_any,omitempty"`

	SysTags *[]TagValuesList `json:"sys_tags,omitempty"`

	Limit *string `json:"limit,omitempty"`

	Offset *string `json:"offset,omitempty"`

	Action string `json:"action"`

	Matches *[]Match `json:"matches,omitempty"`

	WithoutAnyTag *bool `json:"without_any_tag,omitempty"`
}

func (o QueryResourceInstanceTagsBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryResourceInstanceTagsBody struct{}"
	}

	return strings.Join([]string{"QueryResourceInstanceTagsBody", string(data)}, " ")
}
