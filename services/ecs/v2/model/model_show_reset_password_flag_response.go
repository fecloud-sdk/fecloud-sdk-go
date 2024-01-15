package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowResetPasswordFlagResponse struct {
	ResetpwdFlag   *string `json:"resetpwd_flag,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowResetPasswordFlagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResetPasswordFlagResponse struct{}"
	}

	return strings.Join([]string{"ShowResetPasswordFlagResponse", string(data)}, " ")
}
