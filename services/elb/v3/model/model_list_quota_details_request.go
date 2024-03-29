package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListQuotaDetailsRequest struct {
	QuotaKey *[]string `json:"quota_key,omitempty"`
}

func (o ListQuotaDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotaDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListQuotaDetailsRequest", string(data)}, " ")
}
