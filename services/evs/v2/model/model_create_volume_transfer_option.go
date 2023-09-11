package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateVolumeTransferOption struct {

	// 云硬盘过户记录的名称。最大支持255个字节。
	Name string `json:"name"`

	// 云硬盘ID。
	VolumeId string `json:"volume_id"`
}

func (o CreateVolumeTransferOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVolumeTransferOption struct{}"
	}

	return strings.Join([]string{"CreateVolumeTransferOption", string(data)}, " ")
}
