package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowAuditlogPolicyResponse struct {
	KeepDays *int32 `json:"keep_days,omitempty"`

	AuditScope *string `json:"audit_scope,omitempty"`

	AuditTypes     *[]string `json:"audit_types,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowAuditlogPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuditlogPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowAuditlogPolicyResponse", string(data)}, " ")
}
