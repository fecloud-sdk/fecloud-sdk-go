package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowDomainQuotaResponse struct {
	Quotas         *[]DomainQuotaResponseQuotas `json:"quotas,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ShowDomainQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainQuotaResponse", string(data)}, " ")
}
