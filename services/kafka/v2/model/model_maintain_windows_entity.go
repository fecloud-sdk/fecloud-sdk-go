package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type MaintainWindowsEntity struct {
	Default *bool `json:"default,omitempty"`

	End *string `json:"end,omitempty"`

	Begin *string `json:"begin,omitempty"`

	Seq *int32 `json:"seq,omitempty"`
}

func (o MaintainWindowsEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MaintainWindowsEntity struct{}"
	}

	return strings.Join([]string{"MaintainWindowsEntity", string(data)}, " ")
}
