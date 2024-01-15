package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResendReq struct {
	Group *string `json:"group,omitempty"`

	ClientId *string `json:"client_id,omitempty"`

	MsgIdList *[]string `json:"msg_id_list,omitempty"`
}

func (o ResendReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResendReq struct{}"
	}

	return strings.Join([]string{"ResendReq", string(data)}, " ")
}
