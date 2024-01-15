package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowDomainQuotaRequest struct {
	DomainId string `json:"domain_id"`
}

func (o ShowDomainQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainQuotaRequest", string(data)}, " ")
}
