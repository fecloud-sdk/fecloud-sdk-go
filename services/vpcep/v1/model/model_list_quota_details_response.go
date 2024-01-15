package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListQuotaDetailsResponse struct {
	Quotas         *ResourcesResp `json:"quotas,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListQuotaDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotaDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListQuotaDetailsResponse", string(data)}, " ")
}
