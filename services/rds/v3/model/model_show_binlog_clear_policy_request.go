package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowBinlogClearPolicyRequest Request Object
type ShowBinlogClearPolicyRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowBinlogClearPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBinlogClearPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowBinlogClearPolicyRequest", string(data)}, " ")
}
