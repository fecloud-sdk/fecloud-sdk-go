package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SetWtpProtectionStatusInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetWtpProtectionStatusInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetWtpProtectionStatusInfoResponse struct{}"
	}

	return strings.Join([]string{"SetWtpProtectionStatusInfoResponse", string(data)}, " ")
}
