package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AttachServerVolumeRequestBody struct {
	VolumeAttachment *AttachServerVolumeOption `json:"volumeAttachment"`

	DryRun *bool `json:"dry_run,omitempty"`
}

func (o AttachServerVolumeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachServerVolumeRequestBody struct{}"
	}

	return strings.Join([]string{"AttachServerVolumeRequestBody", string(data)}, " ")
}
