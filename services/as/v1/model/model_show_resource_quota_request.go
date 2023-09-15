package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowResourceQuotaRequest Request Object
type ShowResourceQuotaRequest struct {
}

func (o ShowResourceQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceQuotaRequest", string(data)}, " ")
}
