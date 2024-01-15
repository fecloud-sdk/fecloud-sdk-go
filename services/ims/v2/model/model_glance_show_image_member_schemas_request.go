package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type GlanceShowImageMemberSchemasRequest struct {
}

func (o GlanceShowImageMemberSchemasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceShowImageMemberSchemasRequest struct{}"
	}

	return strings.Join([]string{"GlanceShowImageMemberSchemasRequest", string(data)}, " ")
}
