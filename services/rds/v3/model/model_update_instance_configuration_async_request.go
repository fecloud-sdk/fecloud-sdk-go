package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateInstanceConfigurationAsyncRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *UpdateInstanceConfigurationRequestBody `json:"body,omitempty"`
}

func (o UpdateInstanceConfigurationAsyncRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceConfigurationAsyncRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceConfigurationAsyncRequest", string(data)}, " ")
}
