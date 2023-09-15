package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateCustomTemplateRequest Request Object
type CreateCustomTemplateRequest struct {
	Body *CreateCustomTemplateBody `json:"body,omitempty"`
}

func (o CreateCustomTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateCustomTemplateRequest", string(data)}, " ")
}
