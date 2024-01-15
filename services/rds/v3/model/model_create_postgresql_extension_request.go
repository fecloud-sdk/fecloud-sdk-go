package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreatePostgresqlExtensionRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *ExtensionRequest `json:"body,omitempty"`
}

func (o CreatePostgresqlExtensionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostgresqlExtensionRequest struct{}"
	}

	return strings.Join([]string{"CreatePostgresqlExtensionRequest", string(data)}, " ")
}
