package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ExportDlqMessageRequest struct {
	InstanceId string `json:"instance_id"`

	Body *ExportDlqMessageReq `json:"body,omitempty"`
}

func (o ExportDlqMessageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportDlqMessageRequest struct{}"
	}

	return strings.Join([]string{"ExportDlqMessageRequest", string(data)}, " ")
}
