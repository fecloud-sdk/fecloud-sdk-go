package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowInstanceConfigurationRequest Request Object
type ShowInstanceConfigurationRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowInstanceConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceConfigurationRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceConfigurationRequest", string(data)}, " ")
}
