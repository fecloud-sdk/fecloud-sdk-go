package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListProjectTagsResponse struct {
	Tags           *[]QueryProjectTagItem `json:"tags,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListProjectTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectTagsResponse struct{}"
	}

	return strings.Join([]string{"ListProjectTagsResponse", string(data)}, " ")
}
