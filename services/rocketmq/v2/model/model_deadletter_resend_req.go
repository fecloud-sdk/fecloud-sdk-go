package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeadletterResendReq struct {
	Topic *string `json:"topic,omitempty"`

	MsgIdList *[]string `json:"msg_id_list,omitempty"`
}

func (o DeadletterResendReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeadletterResendReq struct{}"
	}

	return strings.Join([]string{"DeadletterResendReq", string(data)}, " ")
}
