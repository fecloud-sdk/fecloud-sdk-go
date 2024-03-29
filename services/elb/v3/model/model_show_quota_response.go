package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowQuotaResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	Quota          *Quota `json:"quota,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowQuotaResponse", string(data)}, " ")
}
