package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DatabaseWithPrivilegeObject struct {
	Name string `json:"name"`

	Readonly *bool `json:"readonly,omitempty"`
}

func (o DatabaseWithPrivilegeObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseWithPrivilegeObject struct{}"
	}

	return strings.Join([]string{"DatabaseWithPrivilegeObject", string(data)}, " ")
}
