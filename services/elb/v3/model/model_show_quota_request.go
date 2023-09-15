package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowQuotaRequest Request Object
type ShowQuotaRequest struct {
}

func (o ShowQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowQuotaRequest", string(data)}, " ")
}
