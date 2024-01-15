package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type Info struct {
	Type *string `json:"type,omitempty"`

	Quota *int64 `json:"quota,omitempty"`

	Used *int64 `json:"used,omitempty"`

	Unit *string `json:"unit,omitempty"`
}

func (o Info) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Info struct{}"
	}

	return strings.Join([]string{"Info", string(data)}, " ")
}
