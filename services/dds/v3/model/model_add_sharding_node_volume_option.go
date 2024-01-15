package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AddShardingNodeVolumeOption struct {
	Size string `json:"size"`
}

func (o AddShardingNodeVolumeOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddShardingNodeVolumeOption struct{}"
	}

	return strings.Join([]string{"AddShardingNodeVolumeOption", string(data)}, " ")
}
