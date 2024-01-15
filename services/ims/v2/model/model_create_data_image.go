package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateDataImage struct {
	Name string `json:"name"`

	VolumeId string `json:"volume_id"`

	Description *string `json:"description,omitempty"`

	Tags *[]string `json:"tags,omitempty"`
}

func (o CreateDataImage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataImage struct{}"
	}

	return strings.Join([]string{"CreateDataImage", string(data)}, " ")
}
