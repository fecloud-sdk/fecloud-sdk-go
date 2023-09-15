package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowQuotasResponse Response Object
type ShowQuotasResponse struct {
	Quotas         *ShowResourcesListResponseBody `json:"quotas,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ShowQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotasResponse struct{}"
	}

	return strings.Join([]string{"ShowQuotasResponse", string(data)}, " ")
}
