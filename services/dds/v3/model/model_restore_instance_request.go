package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RestoreInstanceRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	Body *RestoreInstanceRequestBody `json:"body,omitempty"`
}

func (o RestoreInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreInstanceRequest struct{}"
	}

	return strings.Join([]string{"RestoreInstanceRequest", string(data)}, " ")
}
