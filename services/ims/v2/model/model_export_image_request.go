package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ExportImageRequest Request Object
type ExportImageRequest struct {

	// 镜像ID。
	ImageId string `json:"image_id"`

	Body *ExportImageRequestBody `json:"body,omitempty"`
}

func (o ExportImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportImageRequest struct{}"
	}

	return strings.Join([]string{"ExportImageRequest", string(data)}, " ")
}
