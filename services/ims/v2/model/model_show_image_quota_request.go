package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowImageQuotaRequest Request Object
type ShowImageQuotaRequest struct {
}

func (o ShowImageQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowImageQuotaRequest", string(data)}, " ")
}
