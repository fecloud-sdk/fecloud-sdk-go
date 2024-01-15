package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeadletterResendRespResendResults struct {
	MsgId *string `json:"msg_id,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o DeadletterResendRespResendResults) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeadletterResendRespResendResults struct{}"
	}

	return strings.Join([]string{"DeadletterResendRespResendResults", string(data)}, " ")
}
