package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type GlanceCreateImageMetadataRequest struct {
	Body *GlanceCreateImageMetadataRequestBody `json:"body,omitempty"`
}

func (o GlanceCreateImageMetadataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceCreateImageMetadataRequest struct{}"
	}

	return strings.Join([]string{"GlanceCreateImageMetadataRequest", string(data)}, " ")
}
