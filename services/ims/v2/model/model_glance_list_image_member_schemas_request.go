package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// GlanceListImageMemberSchemasRequest Request Object
type GlanceListImageMemberSchemasRequest struct {
}

func (o GlanceListImageMemberSchemasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceListImageMemberSchemasRequest struct{}"
	}

	return strings.Join([]string{"GlanceListImageMemberSchemasRequest", string(data)}, " ")
}
