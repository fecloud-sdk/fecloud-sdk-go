package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowUserQuotasResponse struct {
	Quotas         *Quotas `json:"quotas,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowUserQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserQuotasResponse struct{}"
	}

	return strings.Join([]string{"ShowUserQuotasResponse", string(data)}, " ")
}
