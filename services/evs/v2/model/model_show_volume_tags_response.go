package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowVolumeTagsResponse Response Object
type ShowVolumeTagsResponse struct {

	// 标签列表。
	Tags           *[]Tag `json:"tags,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowVolumeTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVolumeTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowVolumeTagsResponse", string(data)}, " ")
}
