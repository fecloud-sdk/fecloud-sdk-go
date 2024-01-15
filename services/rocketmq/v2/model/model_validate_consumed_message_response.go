package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ValidateConsumedMessageResponse struct {
	ResendResults  *[]DeadletterResendRespResendResults `json:"resend_results,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o ValidateConsumedMessageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateConsumedMessageResponse struct{}"
	}

	return strings.Join([]string{"ValidateConsumedMessageResponse", string(data)}, " ")
}
