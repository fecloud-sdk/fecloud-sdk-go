package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type GlanceUpdateImageRequest struct {
	ImageId string `json:"image_id"`

	Body *[]GlanceUpdateImageRequestBody `json:"body,omitempty"`
}

func (o GlanceUpdateImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceUpdateImageRequest struct{}"
	}

	return strings.Join([]string{"GlanceUpdateImageRequest", string(data)}, " ")
}
