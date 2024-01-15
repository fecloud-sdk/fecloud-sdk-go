package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type PwdPolicyInfoResponseInfo struct {
	HostId *string `json:"host_id,omitempty"`

	HostName *string `json:"host_name,omitempty"`

	HostIp *string `json:"host_ip,omitempty"`

	MinLength *bool `json:"min_length,omitempty"`

	UppercaseLetter *bool `json:"uppercase_letter,omitempty"`

	LowercaseLetter *bool `json:"lowercase_letter,omitempty"`

	Number *bool `json:"number,omitempty"`

	SpecialCharacter *bool `json:"special_character,omitempty"`

	Suggestion *string `json:"suggestion,omitempty"`
}

func (o PwdPolicyInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PwdPolicyInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"PwdPolicyInfoResponseInfo", string(data)}, " ")
}
