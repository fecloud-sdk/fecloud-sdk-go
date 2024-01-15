package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteGroupReq struct {
	GroupIds *[]string `json:"group_ids,omitempty"`
}

func (o BatchDeleteGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteGroupReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteGroupReq", string(data)}, " ")
}
