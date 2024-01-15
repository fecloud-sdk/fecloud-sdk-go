package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type TemplateId struct {
	TemplateId *int64 `json:"template_id,omitempty"`
}

func (o TemplateId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateId struct{}"
	}

	return strings.Join([]string{"TemplateId", string(data)}, " ")
}
