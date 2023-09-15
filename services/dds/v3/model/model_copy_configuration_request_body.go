package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CopyConfigurationRequestBody struct {

	// 复制后的参数模板模板名称
	Name string `json:"name"`

	// 参数模板模板描述。
	Description *string `json:"description,omitempty"`
}

func (o CopyConfigurationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"CopyConfigurationRequestBody", string(data)}, " ")
}
