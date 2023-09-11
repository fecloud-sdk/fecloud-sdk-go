package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// RestoreExistInstanceRequest Request Object
type RestoreExistInstanceRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	Body *RestoreExistingInstanceRequestBody `json:"body,omitempty"`
}

func (o RestoreExistInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreExistInstanceRequest struct{}"
	}

	return strings.Join([]string{"RestoreExistInstanceRequest", string(data)}, " ")
}
