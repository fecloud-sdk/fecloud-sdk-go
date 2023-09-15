package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowPolicyAndInstanceQuotaResponse Response Object
type ShowPolicyAndInstanceQuotaResponse struct {
	Quotas         *PolicyInstanceQuotas `json:"quotas,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowPolicyAndInstanceQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPolicyAndInstanceQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowPolicyAndInstanceQuotaResponse", string(data)}, " ")
}
