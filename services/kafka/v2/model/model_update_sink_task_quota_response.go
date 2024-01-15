package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateSinkTaskQuotaResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateSinkTaskQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSinkTaskQuotaResponse struct{}"
	}

	return strings.Join([]string{"UpdateSinkTaskQuotaResponse", string(data)}, " ")
}
