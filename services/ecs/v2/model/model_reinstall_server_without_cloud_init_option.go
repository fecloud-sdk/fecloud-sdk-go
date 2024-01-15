package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ReinstallServerWithoutCloudInitOption struct {
	Adminpass *string `json:"adminpass,omitempty"`

	Keyname *string `json:"keyname,omitempty"`

	Userid *string `json:"userid,omitempty"`

	Mode *string `json:"mode,omitempty"`
}

func (o ReinstallServerWithoutCloudInitOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallServerWithoutCloudInitOption struct{}"
	}

	return strings.Join([]string{"ReinstallServerWithoutCloudInitOption", string(data)}, " ")
}
