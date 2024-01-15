package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListAccessPolicyRespPolicies struct {
	AccessKey *string `json:"access_key,omitempty"`

	SecretKey *string `json:"secret_key,omitempty"`

	WhiteRemoteAddress *string `json:"white_remote_address,omitempty"`

	Admin *bool `json:"admin,omitempty"`

	Perm *string `json:"perm,omitempty"`
}

func (o ListAccessPolicyRespPolicies) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessPolicyRespPolicies struct{}"
	}

	return strings.Join([]string{"ListAccessPolicyRespPolicies", string(data)}, " ")
}
