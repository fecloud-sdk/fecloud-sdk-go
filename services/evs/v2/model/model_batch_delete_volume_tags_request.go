package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// BatchDeleteVolumeTagsRequest Request Object
type BatchDeleteVolumeTagsRequest struct {

	// 磁盘ID。
	VolumeId string `json:"volume_id"`

	Body *BatchDeleteVolumeTagsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteVolumeTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteVolumeTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteVolumeTagsRequest", string(data)}, " ")
}
