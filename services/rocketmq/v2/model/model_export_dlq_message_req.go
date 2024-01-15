package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ExportDlqMessageReq struct {
	Topic *string `json:"topic,omitempty"`

	MsgIdList *[]string `json:"msg_id_list,omitempty"`

	UniqKeyList *[]string `json:"uniq_key_list,omitempty"`
}

func (o ExportDlqMessageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportDlqMessageReq struct{}"
	}

	return strings.Join([]string{"ExportDlqMessageReq", string(data)}, " ")
}
