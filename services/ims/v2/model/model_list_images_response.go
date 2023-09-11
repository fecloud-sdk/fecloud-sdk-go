package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ListImagesResponse Response Object
type ListImagesResponse struct {

	// 镜像列表
	Images         *[]ImageInfo `json:"images,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListImagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImagesResponse struct{}"
	}

	return strings.Join([]string{"ListImagesResponse", string(data)}, " ")
}
