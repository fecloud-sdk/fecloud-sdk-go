package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateCustomTemplateResponse Response Object
type CreateCustomTemplateResponse struct {

	// 模板ID
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCustomTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateCustomTemplateResponse", string(data)}, " ")
}
