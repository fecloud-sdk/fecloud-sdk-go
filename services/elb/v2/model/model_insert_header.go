package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type InsertHeader struct {
	XForwardedELBIP *bool `json:"X-Forwarded-ELB-IP,omitempty"`

	XForwardedHost *bool `json:"X-Forwarded-Host,omitempty"`
}

func (o InsertHeader) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InsertHeader struct{}"
	}

	return strings.Join([]string{"InsertHeader", string(data)}, " ")
}
