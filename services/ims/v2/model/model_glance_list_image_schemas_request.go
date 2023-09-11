package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// GlanceListImageSchemasRequest Request Object
type GlanceListImageSchemasRequest struct {
}

func (o GlanceListImageSchemasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceListImageSchemasRequest struct{}"
	}

	return strings.Join([]string{"GlanceListImageSchemasRequest", string(data)}, " ")
}
