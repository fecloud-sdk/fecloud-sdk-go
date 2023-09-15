package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowQuotaOfTenantResponse Response Object
type ShowQuotaOfTenantResponse struct {
	Quotas         *QueryTenantQuotaRespQuotas `json:"quotas,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ShowQuotaOfTenantResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotaOfTenantResponse struct{}"
	}

	return strings.Join([]string{"ShowQuotaOfTenantResponse", string(data)}, " ")
}
