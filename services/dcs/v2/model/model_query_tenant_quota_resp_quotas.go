package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type QueryTenantQuotaRespQuotas struct {
	Resources *[]Resources `json:"resources,omitempty"`
}

func (o QueryTenantQuotaRespQuotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryTenantQuotaRespQuotas struct{}"
	}

	return strings.Join([]string{"QueryTenantQuotaRespQuotas", string(data)}, " ")
}
