package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListVolumeTagsResponse struct {
	Tags           map[string][]string `json:"tags,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListVolumeTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVolumeTagsResponse struct{}"
	}

	return strings.Join([]string{"ListVolumeTagsResponse", string(data)}, " ")
}
