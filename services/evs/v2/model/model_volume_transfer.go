package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type VolumeTransfer struct {
	CreatedAt string `json:"created_at"`

	Id string `json:"id"`

	Links []Link `json:"links"`

	Name string `json:"name"`

	VolumeId string `json:"volume_id"`
}

func (o VolumeTransfer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeTransfer struct{}"
	}

	return strings.Join([]string{"VolumeTransfer", string(data)}, " ")
}
