package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RestoreToExistingInstanceRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	Body *RestoreToExistingInstanceRequestBody `json:"body,omitempty"`
}

func (o RestoreToExistingInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreToExistingInstanceRequest struct{}"
	}

	return strings.Join([]string{"RestoreToExistingInstanceRequest", string(data)}, " ")
}
