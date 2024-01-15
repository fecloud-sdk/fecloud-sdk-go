package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SetAuditlogPolicyRequestBody struct {
	KeepDays int32 `json:"keep_days"`

	ReserveAuditlogs *string `json:"reserve_auditlogs,omitempty"`

	AuditScope *string `json:"audit_scope,omitempty"`

	AuditTypes *[]string `json:"audit_types,omitempty"`
}

func (o SetAuditlogPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAuditlogPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"SetAuditlogPolicyRequestBody", string(data)}, " ")
}
