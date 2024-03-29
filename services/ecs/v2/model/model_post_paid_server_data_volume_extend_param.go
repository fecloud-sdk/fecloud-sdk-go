package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type PostPaidServerDataVolumeExtendParam struct {
	ResourceSpecCode *string `json:"resourceSpecCode,omitempty"`

	ResourceType *string `json:"resourceType,omitempty"`

	SnapshotId *string `json:"snapshotId,omitempty"`
}

func (o PostPaidServerDataVolumeExtendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostPaidServerDataVolumeExtendParam struct{}"
	}

	return strings.Join([]string{"PostPaidServerDataVolumeExtendParam", string(data)}, " ")
}
