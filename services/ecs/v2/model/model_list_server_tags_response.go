package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ListServerTagsResponse Response Object
type ListServerTagsResponse struct {

	// 标签列表
	Tags           *[]ProjectTag `json:"tags,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListServerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerTagsResponse struct{}"
	}

	return strings.Join([]string{"ListServerTagsResponse", string(data)}, " ")
}
