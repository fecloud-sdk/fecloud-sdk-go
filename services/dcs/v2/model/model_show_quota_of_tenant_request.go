package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowQuotaOfTenantRequest Request Object
type ShowQuotaOfTenantRequest struct {
}

func (o ShowQuotaOfTenantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotaOfTenantRequest struct{}"
	}

	return strings.Join([]string{"ShowQuotaOfTenantRequest", string(data)}, " ")
}
