package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowQuotasResponse struct {
	Quotas         *QueryQuotaInfo `json:"quotas,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotasResponse struct{}"
	}

	return strings.Join([]string{"ShowQuotasResponse", string(data)}, " ")
}
