package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// BatchDeleteVolumeTagsResponse Response Object
type BatchDeleteVolumeTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteVolumeTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteVolumeTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteVolumeTagsResponse", string(data)}, " ")
}
