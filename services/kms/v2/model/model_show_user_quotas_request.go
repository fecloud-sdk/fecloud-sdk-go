package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowUserQuotasRequest struct {
}

func (o ShowUserQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserQuotasRequest struct{}"
	}

	return strings.Join([]string{"ShowUserQuotasRequest", string(data)}, " ")
}
