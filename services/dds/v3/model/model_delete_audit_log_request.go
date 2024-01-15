package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteAuditLogRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *DeleteAuditLogRequestBody `json:"body,omitempty"`
}

func (o DeleteAuditLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAuditLogRequest struct{}"
	}

	return strings.Join([]string{"DeleteAuditLogRequest", string(data)}, " ")
}
