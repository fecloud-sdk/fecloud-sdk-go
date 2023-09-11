package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ListImagesTagsRequest Request Object
type ListImagesTagsRequest struct {
}

func (o ListImagesTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImagesTagsRequest struct{}"
	}

	return strings.Join([]string{"ListImagesTagsRequest", string(data)}, " ")
}
