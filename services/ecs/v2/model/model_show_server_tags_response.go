package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowServerTagsResponse Response Object
type ShowServerTagsResponse struct {

	// 标签列表
	Tags           *[]ServerTag `json:"tags,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowServerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowServerTagsResponse", string(data)}, " ")
}
