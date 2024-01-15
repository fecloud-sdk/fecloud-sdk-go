package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SetAuditlogPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetAuditlogPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAuditlogPolicyResponse struct{}"
	}

	return strings.Join([]string{"SetAuditlogPolicyResponse", string(data)}, " ")
}
