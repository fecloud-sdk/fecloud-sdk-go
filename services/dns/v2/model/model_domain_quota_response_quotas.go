package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DomainQuotaResponseQuotas struct {
	QuotaKey string `json:"quota_key"`

	QuotaLimit int32 `json:"quota_limit"`

	Used int32 `json:"used"`

	Unit string `json:"unit"`
}

func (o DomainQuotaResponseQuotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainQuotaResponseQuotas struct{}"
	}

	return strings.Join([]string{"DomainQuotaResponseQuotas", string(data)}, " ")
}
