package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateVolumeTagsRequest struct {
	VolumeId string `json:"volume_id"`

	Body *BatchCreateVolumeTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateVolumeTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateVolumeTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateVolumeTagsRequest", string(data)}, " ")
}
