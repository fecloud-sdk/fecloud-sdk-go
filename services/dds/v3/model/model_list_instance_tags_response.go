package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ListInstanceTagsResponse Response Object
type ListInstanceTagsResponse struct {

	// 标签列表。
	Tags           *[]QueryResourceTagItem `json:"tags,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListInstanceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceTagsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceTagsResponse", string(data)}, " ")
}