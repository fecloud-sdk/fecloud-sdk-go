package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowQuotasRequest Request Object
type ShowQuotasRequest struct {
}

func (o ShowQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotasRequest struct{}"
	}

	return strings.Join([]string{"ShowQuotasRequest", string(data)}, " ")
}
