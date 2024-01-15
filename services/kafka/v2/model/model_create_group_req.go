package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateGroupReq struct {
	GroupName string `json:"group_name"`

	GroupDesc *string `json:"group_desc,omitempty"`
}

func (o CreateGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGroupReq struct{}"
	}

	return strings.Join([]string{"CreateGroupReq", string(data)}, " ")
}
