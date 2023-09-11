package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowReplicationStatusRequest Request Object
type ShowReplicationStatusRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ShowReplicationStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReplicationStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowReplicationStatusRequest", string(data)}, " ")
}
