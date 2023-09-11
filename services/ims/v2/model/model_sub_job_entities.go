package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SubJobEntities struct {

	// 镜像ID。
	ImageId *string `json:"image_id,omitempty"`

	// 镜像名称
	ImageName *string `json:"image_name,omitempty"`
}

func (o SubJobEntities) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubJobEntities struct{}"
	}

	return strings.Join([]string{"SubJobEntities", string(data)}, " ")
}
