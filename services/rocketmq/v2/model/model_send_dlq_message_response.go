package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type SendDlqMessageResponse struct {
	ResendResults  *[]DeadletterResendRespResendResults `json:"resend_results,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o SendDlqMessageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendDlqMessageResponse struct{}"
	}

	return strings.Join([]string{"SendDlqMessageResponse", string(data)}, " ")
}
